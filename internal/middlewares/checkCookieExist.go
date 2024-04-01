package middlewares

import (
	"strings"

	"github.com/Kei-K23/go-otp/internal/config"
	"github.com/Kei-K23/go-otp/internal/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckCookieExist(c *gin.Context) {
	authHeader, err := c.Cookie("go_todo_token")

	if err != nil {
		c.Redirect(303, "/api/v1/login")
		return
	}

	if authHeader == "" {
		c.Redirect(303, "/api/v1/login")
		return
	}

	// Extract the token from the Authorization header
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	tokenString = strings.TrimSpace(tokenString)

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &auth.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.JWT_SECRET_KEY), nil
	})

	// Check for token parsing errors
	if err != nil || !token.Valid {
		c.Redirect(303, "/api/v1/login")
		return
	}

	c.Redirect(303, "/api/v1/users")
	c.Next()
}
