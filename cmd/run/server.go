package run

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/klaus-abram/suncold-restful-app/api/external/owmadapter"
	"github.com/klaus-abram/suncold-restful-app/api/external/storage"
	"github.com/klaus-abram/suncold-restful-app/api/handler"
	"github.com/klaus-abram/suncold-restful-app/api/usecase"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type WeatherServer struct {
	server *http.Server
}

func CreateWeatherServer() *WeatherServer {
	return &WeatherServer{}
}

func (srv *WeatherServer) SunriseWeatherServer(port string, handler http.Handler) error {
	srv.server = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return srv.server.ListenAndServe()
}

func (srv *WeatherServer) SunsetWeatherServer(ctx context.Context) error {
	return srv.server.Shutdown(ctx)
}

func (srv *WeatherServer) RunToShutdownServer(db *sqlx.DB) {

	store := storage.NewStorage(db)
	adapter := owmadapter.NewOwmAdapter()
	cases := usecase.NewUseCase(adapter, store)
	handlers := handler.NewHandler(cases)

	go func() {
		if errInit := srv.SunriseWeatherServer(viper.GetString("port"), handlers.InitWeatherRoutes()); errInit != nil {
			logrus.Fatalf("error occured running http-server %s", errInit.Error())
		}
	}()

	logrus.Print("weather-restful-app - started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("weather-restful-app - shutting down")

	if errShut := srv.SunsetWeatherServer(context.Background()); errShut != nil {
		logrus.Errorf("error occured shutdown http-server %s", errShut.Error())
	}

	if errClose := db.Close(); errClose != nil {
		logrus.Errorf("error with close storage-postgres connection %s", errClose.Error())
	}
}
