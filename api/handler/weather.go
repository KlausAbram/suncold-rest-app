package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get weather in city
// @Security ApiKeyAuth
// @Tags weather
// @Description get weather in city
// @ID get-weeather
// @Produce  json
// @Params dataWeatherResponse body models.WeatherResponse true "weather-info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorJSONResponse
// @Failure 500 {object} errorJSONResponse
// @Failure default {object} errorJSONResponse
// @Router /api/weather [post]

func (hnd *Handler) getWeatherInCity(ctx *gin.Context) {
	agentId, err := getAgentId(ctx)
	if err != nil {
		return
	}

	city := ctx.Param("city")
	if len(city) == 0 {
		newErrorJSONResponse(ctx, http.StatusBadRequest, "location is not found")
		return
	}

	dataWeatherResponse, err := hnd.cases.WeatherSearching.GetWeatherCity(agentId, city)
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, *dataWeatherResponse)
}
