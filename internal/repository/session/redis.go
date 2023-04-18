package session

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	redisPool *redis.Pool
	logger    logging.Logger
}

func NewRepository(redisPool *redis.Pool, logger logging.Logger) Repository {
	return Repository{redisPool: redisPool, logger: logger}
}

func (repo *Repository) GetConnWithContext(ctx context.Context) (redis.ConnWithContext, error) {
	connTmp, err := repo.redisPool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	conn, ok := connTmp.(redis.ConnWithContext)
	if !ok {
		return nil, fmt.Errorf("got connection to radis without context")
	}
	return conn, nil
}

func (repo *Repository) Add(ctx context.Context, session domain.Session) error {
	if session.ExpiresAt.Before(time.Now()) {
		return nil
	}
	sessRow := sessionRow{
		ExpiresAtString: session.ExpiresAt.Format(time.RFC3339),
		UserID:          session.UserID,
	}

	// TODO: сериализую только user_id
	// 1. в session еще есть поля,
	// 2. иммет ли смысл сериализовать UserAgent? (в лекциях его сериализовали)

	dataSerialized, err := json.Marshal(map[string]interface{}{
		"user_id":    sessRow.UserID,
		"expires_at": sessRow.ExpiresAtString,
	})
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return err
	}

	// TODO: session.ID или session.ID.String()
	conn, err := repo.GetConnWithContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	timeToLive := time.Until(session.ExpiresAt)
	result, err := redis.String(conn.DoContext(ctx, "SET", session.ID,
		dataSerialized, "EX", uint(timeToLive.Seconds())))
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Tracef("cant set data in redis: %w", err)
		return err
	}
	if result != "OK" {
		err := fmt.Errorf("'set' in redis replies 'not OK'")
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return err
	}
	return nil
}

func (repo *Repository) Get(ctx context.Context, sessionID uuid.UUID) (domain.Session, error) {
	sessRow := sessionRow{}

	conn, err := repo.GetConnWithContext(ctx)
	if err != nil {
		return domain.Session{}, err
	}
	defer conn.Close()

	data, err := redis.Bytes(conn.DoContext(ctx, "GET", sessionID))
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		if errors.Is(err, redis.ErrNil) {
			return domain.Session{}, domain.ErrSessionNotFound
		}
		return domain.Session{}, fmt.Errorf("cant get data in redis: %w", err)
	}
	err = json.Unmarshal(data, &sessRow)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return domain.Session{}, fmt.Errorf("cant unpack session data from redis: %w", err)
	}
	expireTime, err := time.Parse(time.RFC3339, sessRow.ExpiresAtString)
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return domain.Session{}, nil
	}
	session := domain.Session{
		ExpiresAt: expireTime,
		UserID:    sessRow.UserID,
		ID:        sessionID,
	}
	return session, nil
}

func (repo *Repository) Delete(ctx context.Context, sessionID uuid.UUID) error {
	// TODO: может можно лучше обработать ошибку, зачем приводить к Int? result != OK?
	conn, err := repo.GetConnWithContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = redis.Int(conn.DoContext(ctx, "DEL", sessionID))
	if err != nil {
		repo.logger.WithFields(logrus.Fields{
			"request_id": ctx.Value("requestID").(string),
		}).Trace(err)
		return fmt.Errorf("cant delete by redis: %w", err)
	}
	return nil
}
