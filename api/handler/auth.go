package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/klaus-abram/suncold-restful-app/models"
	"net/http"
)

func (hnd *Handler) SignUp(ctx *gin.Context) {
	var store *models.Agent

	if err := ctx.BindJSON(&store); err != nil {
		newErrorJSONResponse(ctx, http.StatusBadRequest, "invalid request body")
		return
	}

	id, err := hnd.cases.Authorisation.CreateAgent(store)
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (hnd *Handler) SignIn(ctx *gin.Context) {
	var requestData requestSignInData

	if err := ctx.BindJSON(&requestData); err != nil {
		newErrorJSONResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := hnd.cases.Authorisation.CreateJWT(requestData.AgentName, requestData.Password)
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"jwt-token": token,
	})
}
