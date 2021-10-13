package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hnd *Handler) GetCashedRequests(ctx *gin.Context) {
	data, err := hnd.cases.GettingCashedData.GetCashedRequests(ctx)
	if err != nil {
		newErrorJSONResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"requests": data,
	})
}
