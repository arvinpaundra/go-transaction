package tracer

import (
	"clean-arch/pkg/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (*zap.Logger, error) {
	isDevelopment := config.IsDevelopmentEnv()

	cfg := zap.Config{
		Development:      isDevelopment,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			CallerKey:      "caller",
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	return cfg.Build()
}

func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		logger = logger.With(zap.String("method", c.Request.Method))
		logger = logger.With(zap.String("url", c.Request.URL.String()))
		logger = logger.With(zap.String("host", c.Request.Host))
		logger = logger.With(zap.String("ip", c.ClientIP()))

		c.Set("log", logger)

		c.Next()
	}
}

func RecoverMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Fatal("Panic Recovered", zap.Any("error", err))
			}
		}()

		c.Next()
	}
}

func Log(c *gin.Context, key string, value interface{}) {
	logger, ok := c.Get("log")
	if !ok {
		return
	}

	logger.(*zap.Logger).With(zap.Any(key, value))

	c.Set("log", logger)
}
