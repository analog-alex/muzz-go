package middleware

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/types"
	"muzz-service/pkg/types/cryptography"
	"net/http"
)

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			types.ErrResp(c, http.StatusUnauthorized, "token required", nil)
			c.Abort()
			return
		}

		// assert token is Bearer
		if len(token) < 7 || token[:7] != "Bearer " {
			types.ErrResp(
				c, http.StatusUnauthorized, "invalid token format", &types.Extras{"invalid_token": token})
			c.Abort()
			return
		}

		// remove Bearer from token and validate it
		userId, err := cryptography.VerifyJWToken(token[7:])
		if err != nil {
			types.ErrResp(
				c, http.StatusUnauthorized, "invalid token", &types.Extras{"invalid_token": token})
			c.Abort()
			return
		}

		// set context and continue
		c.Set("userId", userId) // we should enrich this with more user data
		c.Next()
	}
}
