package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/test_env"
	"github.com/stretchr/testify/require"
)

var testCasesWithCookie = []test_env.TestCase{
	{
		// signup с куки: запрещено
		Path:   "/signup",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    test_env.TestUser.Email,
			"password": test_env.TestUser.Password,
		},
		WithCookie: true,
		StatusCode: 403,
	},
	{
		// signin с куки: запрещено
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    test_env.TestUser.Email,
			"password": test_env.TestUser.Password,
		},
		WithCookie: true,
		StatusCode: 403,
	},
	{
		// profile с куки: разрешено
		Path:       "/profile",
		Method:     "GET",
		WithCookie: true,
		StatusCode: 200,
	},
	{
		// для selections кука опциональна
		Path:       "/selections",
		Method:     "GET",
		WithCookie: true,
		StatusCode: 200,
	},
	{
		// для selections кука опциональна
		Path:       "/selections",
		Method:     "GET",
		WithCookie: false,
		StatusCode: 200,
	},
	{
		// profile без куки: запрещено
		Path:       "/profile",
		Method:     "GET",
		WithCookie: false,
		StatusCode: 401,
	},
	{
		// logout без куки: запрещено
		Path:       "/logout",
		Method:     "POST",
		WithCookie: false,
		StatusCode: 401,
	},
	{
		// logout с куки: разрешено
		Path:       "/logout",
		Method:     "POST",
		WithCookie: true,
		StatusCode: 200,
	},
	{
		// profile с "протухшей" кукой: запрещено
		Path:       "/profile",
		Method:     "GET",
		WithCookie: true,
		StatusCode: 401,
	},
	{
		// signin с "протухшей" кукой: на сервере она удалена, поэтому сервер просто выдаст новую
		Path:   "/signin",
		Method: "POST",
		RequestBody: map[string]interface{}{
			"email":    test_env.TestUser.Email,
			"password": test_env.TestUser.Password,
		},
		WithCookie: true,
		StatusCode: 200,
	},
}

func TestApiCookie(t *testing.T) {
	testEnv := test_env.NewTestEnv()
	// регистрация
	reqBody, err := json.Marshal(test_env.TestUser)
	if err != nil {
		t.Errorf("internal error: error while unmarshalling JSON: %s", err)
	}
	reqBodyReader := bytes.NewReader(reqBody)
	req := httptest.NewRequest("POST", "/signup", reqBodyReader)
	w := httptest.NewRecorder()
	testEnv.Router.ServeHTTP(w, req)
	require.Equal(t, http.StatusCreated, w.Code, fmt.Sprintf("TestApiCookie signup, wrong status %d, expected %d", w.Code, http.StatusCreated))

	// авторизация

	reqBodyReader = bytes.NewReader(reqBody)
	req = httptest.NewRequest("POST", "/signin", reqBodyReader)
	w = httptest.NewRecorder()
	testEnv.Router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code, fmt.Sprintf("TestApiCookie signin, wrong status %d, expected %d", w.Code, http.StatusOK))
	cookie := w.Result().Header.Get("Set-Cookie")

	for numCase, testCase := range testCasesWithCookie {
		var reqBodyReader io.Reader // reqBodyReader = nil
		if testCase.RequestBody != nil {
			reqBodyReader = bytes.NewReader(reqBody)
		}
		req = httptest.NewRequest(testCase.Method, testCase.Path, reqBodyReader)
		if testCase.WithCookie {
			req.Header.Add("Cookie", cookie)
		}
		w = httptest.NewRecorder()
		testEnv.Router.ServeHTTP(w, req)
		require.Equal(t, testCase.StatusCode, w.Code, fmt.Sprintf("TestApiCookie [%d], wrong status %d, expected %d", numCase, w.Code, testCase.StatusCode))
	}
}
