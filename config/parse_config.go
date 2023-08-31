package config

import (
	"rest/internal/models"

	"github.com/spf13/viper"
)

func ParseConfig() (*models.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &models.Config{
		HostRedis:   viper.GetString("redis.host"),
		PortRedis:   viper.GetString("redis.port"),
		HostSQL:     viper.GetString("postgres.host"),
		PortSQL:     viper.GetString("postgres.port"),
		UsernameSQL: viper.GetString("postgres.username"),
		PasswordSQL: viper.GetString("postgres.password"),
		DBName:      viper.GetString("postgres.dbname"),
		SSLmode:     viper.GetString("postgres.sslmode"),
	}, nil
}
