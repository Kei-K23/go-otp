package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ContextKey string

const ClaimsContextKey ContextKey = "claims"

func WriteJSON(c *gin.Context, status int, payload any) {
	c.Status(status)
	c.JSON(status, payload)
}

func WriteError(c *gin.Context, status int, errPayload any) {
	WriteJSON(c, status, errPayload)
}

func GetCookieValue(c *gin.Context, name string) (string, error) {
	value, exists := c.Get(string(ClaimsContextKey))
	if !exists {

		return "", fmt.Errorf("missing cookie for name %s", name)
	}

	// Convert userID to the appropriate type
	v, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("cookie value is not of type string")
	}
	return v, nil
}
