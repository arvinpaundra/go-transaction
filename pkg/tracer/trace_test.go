package tracer_test

import (
	"clean-arch/pkg/tracer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitLogger(t *testing.T) {
	logger, err := tracer.InitLogger()
	assert.Nil(t, err)
	assert.NotNil(t, logger)
}

func TestLoggingMiddleware(t *testing.T) {
	logger, err := tracer.InitLogger()
	assert.Nil(t, err)
	assert.NotNil(t, logger)

	middleware := tracer.LoggingMiddleware(logger)
	assert.NotNil(t, middleware)
}

func TestRecoverMiddleware(t *testing.T) {
	logger, err := tracer.InitLogger()
	assert.Nil(t, err)
	assert.NotNil(t, logger)

	middleware := tracer.RecoverMiddleware(logger)
	assert.NotNil(t, middleware)
}
