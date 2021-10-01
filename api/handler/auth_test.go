package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/klaus-abram/suncold-restful-app/api/usecase"
	casemocks "github.com/klaus-abram/suncold-restful-app/api/usecase/mocks"
	"github.com/klaus-abram/suncold-restful-app/models"
	//"github.com/stretchr/testify/mock"
)

func TestHandler_signUp(t *testing.T) {

	//mock behavior
	type mockBehavior func(r *casemocks.MockAuthorisation, agent models.Agent)

	//test table
	testTable := []struct {
		name                 string
		requestBody          string
		agentBody            models.Agent
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "CORRECT",
			requestBody: `{"name": "name", "agent_name": "name", "password": "password"}`,
			mockBehavior: func(r *casemocks.MockAuthorisation, agent models.Agent) {
				r.EXPECT().CreateAgent(agent).Return(1, nil)
			},
			agentBody: models.Agent{
				Name:      "name",
				AgentName: "name",
				Password:  "password"},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:                 "INVALID REQUEST DATA",
			requestBody:          `{"agent_name": "username"}`,
			agentBody:            models.Agent{},
			mockBehavior:         func(r *casemocks.MockAuthorisation, agent models.Agent) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid request data"}`,
		},
		{
			name:        "INTERNAL ERROR",
			requestBody: `{"name": "name", "agent_name": "name", "password": "password"}`,
			mockBehavior: func(r *casemocks.MockAuthorisation, agent models.Agent) {
				r.EXPECT().CreateAgent(agent).Return(0, errors.New("internal server error"))
			},
			agentBody: models.Agent{
				Name:      "name",
				AgentName: "name",
				Password:  "password"},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"internal server error"}`,
		},
	}

	//act
	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			//init dep
			c := gomock.NewController(t)
			defer c.Finish()

			store := casemocks.NewMockAuthorisation(c)
			test.mockBehavior(store, test.agentBody)

			ucase := &usecase.UseCase{Authorisation: store}
			handler := Handler{ucase}

			//init Endpoint
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			//init Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(test.requestBody))

			//make Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}

}
