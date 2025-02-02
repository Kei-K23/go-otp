package middlewares

import (
	"strings"

	"github.com/Kei-K23/go-otp/internal/config"
	"github.com/Kei-K23/go-otp/internal/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const ClaimsContextKey ContextKey = "claims"

func AuthMiddleware(c *gin.Context) {
	authHeader, err := c.Cookie("go_todo_token")

	if err != nil {
		c.Redirect(303, "/api/v1/login")
		return
	}

	// Check if Authorization header is present
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

	// Extract claims from the token
	claims, ok := token.Claims.(*auth.JWTClaim)
	if !ok {
		c.Redirect(303, "/api/v1/login")
		return
	}

	// Set claims in the context
	c.Set(string(ClaimsContextKey), claims.UserID)
	c.Next()
}
