package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/klaus-abram/suncold-restful-app/api/usecase"
)

type Handler struct {
	cases *usecase.UseCase
}

func NewHandler(cases *usecase.UseCase) *Handler {
	return &Handler{cases: cases}
}

func (hnd *Handler) InitWeatherRoutes() *gin.Engine {
	weatherRouter := gin.New()

	authWeather := weatherRouter.Group("/auth")
	{
		authWeather.POST("/sing-up", hnd.SignUp)
		authWeather.POST("/sing-in", hnd.SignIn)
	}

	getWeather := weatherRouter.Group("/weather", hnd.agentIdentity)
	{
		getWeather.POST("/:city", hnd.getWeatherInCity)

		historyWeather := getWeather.Group("/history")
		{
			historyWeather.GET("/location/:city", hnd.GetAllHistoryCity)
			historyWeather.GET("/moment/:day", hnd.GetAllHistoryMoment)
		}

	}

	weatherRouter.GET("/requests/history/:agent", hnd.GetAgentHistory)

	return weatherRouter
}
