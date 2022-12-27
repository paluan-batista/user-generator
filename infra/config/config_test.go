package config

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitLoadWithEnvironmentVariablesByDefault(t *testing.T) {
	config := Init() //LOCALHOST
	assert.Equal(t, config.Server.ServiceName, "user-generator")
	assert.Equal(t, config.Server.Host, "0.0.0.0:8000")
	assert.Equal(t, config.Server.MetaHost, "0.0.0.0:8001")
	assert.Equal(t, config.Database.MySqlConnection, "register_user:register_pwd@tcp(localhost:3308)/register_db?charset=utf8mb4&parseTime=True&loc=UTC")
}

func TestInitLoadWithEnvironmentVariables(t *testing.T) {
	err := godotenv.Load(".env") //QA and PRD
	assert.NoError(t, err)

	config := Init() //LOCALHOST
	assert.Equal(t, config.Server.ServiceName, "user-generator")
	assert.Equal(t, config.Server.Host, "0.0.0.0:8000")
	assert.Equal(t, config.Server.MetaHost, "0.0.0.0:8001")
	assert.Equal(t, config.Database.MySqlConnection, "register_user:register_pwd@tcp(localhost:3308)/register_db?charset=utf8mb4&parseTime=True&loc=UTC")

}
