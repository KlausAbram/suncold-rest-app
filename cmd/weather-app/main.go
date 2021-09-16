package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/klaus-abram/suncold-restful-app/cmd/run"
	"github.com/klaus-abram/suncold-restful-app/migrations"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {

	if err := run.SetLoggingConfig(); err != nil {
		logrus.Fatalf("error with setting configs and format: [%s]", err.Error())
	}

	db, err := run.InitPostgresStorage()
	if err != nil {
		logrus.Fatalf("error with initialise db/db connection: [%s]", err.Error())
	}

	if err := migrations.UpMigrations(db.DB); err != nil {
		logrus.Fatalf("error with init or up-migration - [%s]", err.Error())
	}

	//create server-object
	serv := run.CreateWeatherServer()

	//run server and shutdown in time it needs
	serv.RunToShutdownServer(db)
}
