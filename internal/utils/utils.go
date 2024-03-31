package utils

import (
	"github.com/gin-gonic/gin"
)

func WriteJSON(c *gin.Context, status int, payload any) {
	c.Status(status)
	c.JSON(status, payload)
}

func WriteError(c *gin.Context, status int, errPayload any) {
	WriteJSON(c, status, errPayload)
}
