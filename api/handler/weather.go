package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (hnd *Handler) getWeatherInCity(ctx *gin.Context) {
	agentId, err := getAgentId(ctx)
	if err != nil {
		return
	}

	city := ctx.Param("city")
	if len(city) == 0 {
		newErrorJSONResponse(ctx, http.StatusBadRequest, "city is not found")
		return
	}

	data, err := hnd.cases.WeatherSearching.GetWeatherCity(agentId, city)
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, data)
}
