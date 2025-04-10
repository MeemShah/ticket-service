package config

import (
	"log/slog"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() error {

	exit := func(err error) {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		slog.Warn(".env not found, that's okay!")
	}

	viper.AutomaticEnv()

	config = &Config{
		Mode:        Mode(viper.GetString("MODE")),
		ServiceName: viper.GetString("SERVICE_NAME"),
		HttpPort:    viper.GetInt("HTTP_PORT"),
		JwtSecret:   viper.GetString("JWT_SECREAT"),

		ReadRedisUrl:           viper.GetString("READ_REDIS_URL"),
		WriteRedisUrl:          viper.GetString("WRITE_REDIS_URL"),
		EnableSSLMode:          viper.GetBool("ENABLE_SSL_MODE"),
		HoldTicktViewInSeconds: viper.GetInt("HOLD_TICKET_VIEW_IN_SECONDS"),

		//MigrationSource:   viper.GetString("MIGRATION_SOURCE"),
		// DB: &DBConfig{
		// 	Host:                 viper.GetString("DB_HOST"),
		// 	Port:                 viper.GetInt("DB_PORT"),
		// 	Name:                 viper.GetString("DB_NAME"),
		// 	User:                 viper.GetString("DB_USER"),
		// 	Password:             viper.GetString("DB_PASS"),
		// 	MaxIdleTimeInMinutes: viper.GetInt("MAX_IDLE_TIME_IN_MINUTE"),
		// 	EnableSSLMode:        viper.GetBool("ENABLE_SSL_MODE"),
		// },
	}

	v := validator.New()
	if err = v.Struct(config); err != nil {
		exit(err)
	}

	return nil
}
