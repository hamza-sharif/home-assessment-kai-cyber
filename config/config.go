// Package config provides configuration settings for the application.
package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration.
const (
	ServerHost = "server.host"
	ServerPort = "server.port"

	GitRetries       = "git.retries"
	ConcurrencyLimit = "concurrency.limit"

	LogLevel = "log.level"
)

func init() {
	// env var for db.
	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	_ = viper.BindEnv(GitRetries, "GIT_RETRIES")
	_ = viper.BindEnv(ConcurrencyLimit, "CONCURRENCY_LIMIT")

	_ = viper.BindEnv(LogLevel, "LOG_LEVEL")

	// defaults.
	defaultVariableValues()
}

func defaultVariableValues() {
	viper.SetDefault(ServerHost, "0.0.0.0")
	viper.SetDefault(ServerPort, "8080")

	viper.SetDefault(GitRetries, "3")
	viper.SetDefault(ConcurrencyLimit, 3)

	viper.SetDefault(LogLevel, "debug")

}
