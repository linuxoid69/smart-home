package config

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

func GetConfig() {
	viper.SetConfigFile(os.Getenv("CONFIG"))

	if err := viper.ReadInConfig(); err != nil {
		slog.Error("Error reading config file:", "error", err)
	}
}
