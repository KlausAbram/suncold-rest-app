package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klaus-abram/suncold-restful-app/models"
)

// @Summary SignUp
// @Tags auth
// @Description sign-up agent
// @ID create-account
// @Accept  json
// @Produce  json
// @Param agent body models.Agent true "agent-info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorJSONResponse
// @Failure 500 {object} errorJSONResponse
// @Failure default {object} errorJSONResponse
// @Router /auth/sign-up [post]

func (hnd *Handler) signUp(ctx *gin.Context) {
	var agent models.Agent

	if err := ctx.BindJSON(&agent); err != nil {
		newErrorJSONResponse(ctx, http.StatusBadRequest, "invalid request data")
		return
	}

	id, err := hnd.cases.Authorisation.CreateAgent(agent)
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary SignIn
// @Tags auth
// @Description authentication
// @ID login
// @Accept  json
// @Produce  json
// @Param requestData body handler.equestSignInData true "credentials"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorJSONResponse
// @Failure 500 {object} errorJSONResponse
// @Failure default {object} errorJSONResponse
// @Router /auth/sign-in [post]

func (hnd *Handler) signIn(ctx *gin.Context) {
	var requestData requestSignInData

	if err := ctx.BindJSON(&requestData); err != nil {
		newErrorJSONResponse(ctx, http.StatusBadRequest, "invalid request data")
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
