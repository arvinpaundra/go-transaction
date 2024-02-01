package middleware

import (
	"clean-arch/pkg/consts"
	"clean-arch/pkg/crypto"
	"clean-arch/pkg/util"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
AuthorizeSignature is middleware for internal service communication

@param gin.Context

@return gin.HandlerFunc
*/
func AuthorizeSignature() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			secretKey = util.GetEnv("APP_SECRET_KEY", "komsecret")
			authToken = c.GetHeader("Authorization")
			dateTime  = c.GetHeader("Date-Time")
			signature = c.GetHeader("Signature")
			tNow      = time.Now().AddDate(0, 0, -2) // time for 2 days ago
		)

		err := func() error {
			dateTimeTime, err := time.Parse(consts.TimeFormatDateTime, dateTime)
			if err != nil {
				return errors.New("failed to parse date-time, please use format yyyy-mm-dd hh:mm:ss")
			}

			// check if signature was expired (2 days ago)
			if dateTimeTime.Before(tNow) {
				return errors.New("signature was expired")
			}

			generatedSignature := crypto.EncodeSHA256HMAC(secretKey, authToken, dateTime)
			if generatedSignature != signature {
				return errors.New("signature not match")
			}

			return nil
		}()
		if err != nil {
			log.Println("unauthorize signature: ", err)

			response := util.APIResponse("Unauthorized", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

		c.Next()
	}
}
