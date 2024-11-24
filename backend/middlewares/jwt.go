package middlewares

import (
	errors2 "Inkspire/errors"
	"Inkspire/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthMiddleware is a middleware that checks for a valid JWT token.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, errors2.ErrUserAuthFailed)
			c.Abort()
			return
		}

		// Extract the token from the header
		tokenString := authHeader[len("Bearer "):]
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, errors2.ErrUserAuthFailed)
			c.Abort()
			return
		}

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				c.JSON(http.StatusUnauthorized, errors2.ErrUserAuthFailed)
			} else {
				c.JSON(http.StatusUnauthorized, errors2.ErrUserAuthFailed)
			}
			c.Abort()
			return
		}

		// Set the username in the context for use in subsequent handlers
		c.Set("username", claims.Username)
		c.Next()
	}
}
