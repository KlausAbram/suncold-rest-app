package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/klaus-abram/suncold-restful-app/api/usecase"
	casemocks "github.com/klaus-abram/suncold-restful-app/api/usecase/mocks"
	mock_usecase "github.com/klaus-abram/suncold-restful-app/api/usecase/mocks"
)

func TestHandler_agentIdentity(t *testing.T) {
	type mockBehavior func(r *casemocks.MockAuthorisation, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Correct",
			headerName:  "Authorization",
			headerValue: "Bearer test",
			token:       "test",
			mockBehavior: func(r *casemocks.MockAuthorisation, token string) {
				r.EXPECT().ParseJWT(token).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "Invalid header name",
			headerName:           "",
			headerValue:          "Bearer test",
			token:                "test",
			mockBehavior:         func(r *casemocks.MockAuthorisation, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Invalid header value",
			headerName:           "Authorization",
			headerValue:          "Beareasfr test",
			token:                "test",
			mockBehavior:         func(r *casemocks.MockAuthorisation, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:                 "Empty token",
			headerName:           "Authorization",
			headerValue:          "Bearer",
			token:                "",
			mockBehavior:         func(r *casemocks.MockAuthorisation, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name:        "Parse JWT-token error",
			headerName:  "Authorization",
			headerValue: "Bearer test",
			token:       "test",
			mockBehavior: func(r *casemocks.MockAuthorisation, token string) {
				r.EXPECT().ParseJWT(token).Return(0, errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {

			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_usecase.NewMockAuthorisation(c)
			test.mockBehavior(repo, test.token)

			ucase := &usecase.UseCase{Authorisation: repo}
			handler := Handler{cases: ucase}

			r := gin.New()
			r.GET("/identity", handler.agentIdentity, func(c *gin.Context) {
				id, _ := c.Get(agentCtx)
				c.String(200, "%d", id)
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/identity", nil)
			req.Header.Set(test.headerName, test.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
