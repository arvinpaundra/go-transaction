package middleware

import (
	"clean-arch/internal/dto"
	"clean-arch/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Can be used in Http package or in every router.go inside each service
func ApiKeyAuth() gin.HandlerFunc {
	apiKey := util.GetEnv("API_KEY", "fallback")
	return func(c *gin.Context) {
		inputKey := c.Request.Header["Authorization"]
		if len(inputKey) == 0 || inputKey[0][7:] != apiKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Common{
				Status:  "failed",
				Code:    401,
				Message: "Unautenticated",
			})
		} else {
			c.Next()
		}
	}
}