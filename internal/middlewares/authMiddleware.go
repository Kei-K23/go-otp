package middlewares

import (
	"net/http"
	"strings"

	"github.com/Kei-K23/go-otp/internal/config"
	"github.com/Kei-K23/go-otp/internal/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const ClaimsContextKey ContextKey = "claims"

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	// Check if Authorization header is present
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	// Extract claims from the token
	claims, ok := token.Claims.(*auth.JWTClaim)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unable to extract claims from token"})
		return
	}

	// Set claims in the context
	c.Set(string(ClaimsContextKey), claims.UserID)

	c.Next()
}
