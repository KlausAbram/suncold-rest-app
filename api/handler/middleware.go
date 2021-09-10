package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type requestSignInData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type errorJSONResponse struct {
	Message string `json:"message"`
}

func newErrorJSONResponse(ctx *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorJSONResponse{Message: message})
}

func (hnd *Handler) userIdentity(ctx *gin.Context) {

}
