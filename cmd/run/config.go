package run

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// SetLoggingConfig SetVariableLoggingConfig Set configs/env and logging-mode
// make in main
func SetLoggingConfig() error {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := SetViperConfig(); err != nil {
		return err
	}

	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}

// SetViperConfig Set viper-config
func SetViperConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
