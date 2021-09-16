package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	authHeader = "Authorisation"
	agentCtx   = "agentId"
)

type requestSignInData struct {
	AgentName string `json:"agent_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type errorJSONResponse struct {
	Message string `json:"message"`
}

func newErrorJSONResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorJSONResponse{Message: message})
}

func (hnd *Handler) agentIdentity(ctx *gin.Context) {
	authHead := ctx.GetHeader(authHeader)

	headerSections := strings.Split(authHead, " ")
	if authHead != "" || len(headerSections) != 2 || headerSections[0] != "Bearer" {
		newErrorJSONResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerSections[1]) == 0 {
		newErrorJSONResponse(ctx, http.StatusUnauthorized, "token is empty")
		return
	}

	agentId, err := hnd.cases.Authorisation.ParseJWT(headerSections[1])
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(agentCtx, agentId)
}

func getAgentId(ctx *gin.Context) (int, error) {
	id, ok := ctx.Get(agentCtx)
	if !ok {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, "agent id not found")
		return 0, nil
	}

	agentId, ok := id.(int)
	if !ok {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, "agent id isn't of valid type")
		return 0, nil
	}

	return agentId, nil
}
