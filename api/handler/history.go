package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (hnd *Handler) GetAllHistoryCity(ctx *gin.Context) {
	location := ctx.Param("location")
	if len(location) == 0 {
		newErrorJSONResponse(ctx, http.StatusBadRequest, "location is not found")
		return
	}

	dataStatesResponse, err := hnd.cases.GettingWeatherHistory.GetHistoryLocation(location)
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, dataStatesResponse)
}

func (hnd *Handler) GetAllHistoryMoment(ctx *gin.Context) {
	moment := ctx.Param("moment")
	if _, err := time.Parse("2006-01-02", moment); err != nil {
		newErrorJSONResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	dataRequests, err := hnd.cases.GettingWeatherHistory.GetHistoryMoment(moment)
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, dataRequests)
}

func (hnd *Handler) getAgentHistory(ctx *gin.Context) {
	agent := ctx.Param("agent")
	if len(agent) == 0 {
		newErrorJSONResponse(ctx, http.StatusBadRequest, "name is empty")
		return
	}

	agentRequestsData, err := hnd.cases.GettingWeatherHistory.GetAgentHistory(agent)
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, agentRequestsData)
}
