package config_test

import (
	"clean-arch/pkg/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMysqlHost(t *testing.T) {
	result := config.GetMysqlHost()
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}

func TestGetMysqlPassword(t *testing.T) {
	result := config.GetMysqlPassword()
	assert.NotNil(t, result)
}

func TestGetMysqlPort(t *testing.T) {
	result := config.GetMysqlPort()
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}

func TestGetMysqlUser(t *testing.T) {
	result := config.GetMysqlUser()
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}

func TestGetEnvironment(t *testing.T) {
	result := config.GetEnvironment()
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}
