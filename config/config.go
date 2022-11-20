package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

func InitConfig(conf *viper.Viper, logger zap.Logger) {
	conf.SetDefault("PORT", "8080")
	conf.SetDefault("ENVIRONMENT", "DEVELOPMENT")
	conf.SetDefault("REDIS_HOST", "redis")
	conf.SetDefault("REDIS_PORT", "6379")
	conf.SetDefault("REDIS_PASSWORD", "")
	conf.SetDefault("POLLING_RATES_INTERVAL", 50*time.Second)

	conf.AddConfigPath(".")
	conf.AddConfigPath("/app")
	conf.SetConfigFile("config.yaml")
	conf.SetConfigFile(".env")
	conf.SetConfigFile("/app/.env")

	if err := conf.ReadInConfig(); err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logger.Info("Using default settings, config file not found", zap.Error(err))
		} else {
			// Config file was found but another error was produced
			logger.Info("Config file was found, but an error occurred", zap.Error(err))
		}
	}

	conf.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("Config file changed", zap.String("File", e.Name))
	})
	conf.WatchConfig()
}
