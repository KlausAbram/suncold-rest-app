package main

import (
	"github.com/klaus-abram/suncold-restful-app/cmd/init"
	"github.com/sirupsen/logrus"
)

// docker run --name=weather-db -e POSTGRES_PASSWORD=klaus -p 5436:5432 -d --rm postgres
// migrate -path ./schema -database postgres://postgres:klaus@localhost:5436/postgres?sslmode=disable up
func main() {

	if err := init.SetLoggingConfig(); err != nil {
		logrus.Fatalf("error with setting configs and format: [%s]", err.Error())
	}

	db, err := init.InitPostgresStorage()
	if err != nil {
		logrus.Fatalf("error with initialise db/db connection: [%s]", err.Error())
	}

	//create server-object
	serv := init.CreateWeatherServer()

	//run server and shutdown in time it needs
	serv.RunToShutdownServer(db)

}
