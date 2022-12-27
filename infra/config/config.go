package config

import (
	"github.com/spf13/viper"
	"os"
	"user-generator/infra/log"
)

var (
	config    = Init()
	logConfig = log.NewLogger()
)

type Config struct {
	Server   *ServerConfig   `mapstructure:"SERVER_CONFIG"`
	Database *DatabaseConfig `mapstructure:"DATABASE"`
}

type ServerConfig struct {
	Host        string `mapstructure:"SERVER_HOST"`
	MetaHost    string `mapstructure:"META_HOST"`
	ServiceName string `mapstructure:"SERVICE_NAME"`
}

type DatabaseConfig struct {
	MySqlConnection string `mapstructure:"MYSQL_CONNECTION"`
}

func Init() *Config {

	viper.SetDefault("SERVICE_NAME", getEnvToString("SERVICE_NAME", "user-generator"))
	viper.SetDefault("SERVER_HOST", getEnvToString("SERVER_HOST", "0.0.0.0:8000"))
	viper.SetDefault("META_HOST", getEnvToString("META_HOST", "0.0.0.0:8001"))
	viper.SetDefault("MYSQL_CONNECTION", getEnvToString("MYSQL_CONNECTION", "register_user:register_pwd@tcp(localhost:3308)/register_db?charset=utf8mb4&parseTime=True&loc=UTC"))

	appConfig := &Config{}
	if err := viper.Unmarshal(&appConfig); err != nil {
		logConfig.Fatal(err.Error())
	}

	serverConfig := &ServerConfig{}
	if err := viper.Unmarshal(&serverConfig); err != nil {
		logConfig.Fatal(err.Error())
	}

	databaseConfig := &DatabaseConfig{}
	if err := viper.Unmarshal(&databaseConfig); err != nil {
		logConfig.Fatal(err.Error())
	}

	appConfig.Server = serverConfig
	appConfig.Database = databaseConfig
	return appConfig

}

func GetServerConfig() *ServerConfig {
	return config.Server
}

func GetDatabaseConfig() *DatabaseConfig {
	return config.Database
}

func getEnvToString(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
