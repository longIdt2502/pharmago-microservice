package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER_ACCOUNT   string        `mapstructure:"DB_DRIVER_ACCOUNT"`
	DB_SOURCE_ACCOUNT   string        `mapstructure:"DB_SOURCE_ACCOUNT"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress   string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	RedisAddress        string        `mapstructure:"REDIS_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	EmailSenderName     string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	B2KeyId             string        `mapstructure:"B2_KEY_ID"`
	B2KeyName           string        `mapstructure:"B2_KEY_NAME"`
	B2ApplicationKey    string        `mapstructure:"B2_APPLICATION_KEY"`
	B2Bucket            string        `mapstructure:"B2_BUCKET"`
	B2AccountId         string        `mapstructure:"B2_ACCOUNT_ID"`
	MigrationURL        string        `mapstructure:"MIGRATION_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&config)
	return
}
