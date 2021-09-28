package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (hnd *Handler) GetAgentHistory(ctx *gin.Context) {

}

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

	logrus.Print(dataStatesResponse)

	ctx.JSON(http.StatusOK, dataStatesResponse)
}

func (hnd *Handler) GetAllHistoryMoment(ctx *gin.Context) {

}
