package setup

import (
	"context"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type SessionUseCase interface {
	Create(ctx context.Context, user domain.User) (domain.Session, error)
	Get(ctx context.Context, sessionID string) (domain.Session, error)
	Delete(ctx context.Context, sessionID string) error
}
