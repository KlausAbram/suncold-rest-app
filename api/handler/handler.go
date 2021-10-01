package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/klaus-abram/suncold-restful-app/api/usecase"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
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
		authWeather.POST("/sign-up", hnd.signUp)
		authWeather.POST("/sign-in", hnd.signIn)
	}

	api := weatherRouter.Group("/api", hnd.agentIdentity)
	{

		getWeather := api.Group("/weather")
		{
			getWeather.POST("/:city", hnd.getWeatherInCity)

		}

		getHistory := api.Group("/history")
		{
			getHistory.GET("/location/:location", hnd.GetAllHistoryCity)
			getHistory.GET("/moment/:moment", hnd.GetAllHistoryMoment)
		}

		api.GET("/forecast/:location", hnd.getForecast)
		api.GET("/requests/:agent", hnd.getAgentHistory)
	}

	//weatherRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return weatherRouter
}
