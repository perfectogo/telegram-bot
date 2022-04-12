package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type Config struct {
	ChatUsername string
	ServicePORT  string
	ServiceHost  string
	ClientPORT   string
	LogLevel     string
}

var ChatUsername string
var TelegramBotToken string

// Load loads environment vars and inflates Config
func Load() (Config, error) {
	cfg := Config{}

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	ChatUsername = viper.GetString("tg.chatusername")
	TelegramBotToken = os.Getenv("TG_API_TOKEN")
	cfg.ServiceHost = cast.ToString(getOrReturnDefault("CATALOG_SERVICE_HOST", "127.0.0.1"))
	cfg.ServicePORT = viper.GetString("ports.servicePort")
	cfg.ClientPORT = viper.GetString("ports.clientPort")
	cfg.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	return cfg, nil
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
