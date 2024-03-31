package auth

import (
	"net/http"

	"github.com/Kei-K23/go-otp/internal/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authStore types.AuthStore
}

type Login struct {
	Username string `json:"username" binding:"required"`
}

func NewHandler(authStore types.AuthStore) *Handler {
	return &Handler{authStore: authStore}
}

func (h *Handler) RegisterRoutes(router gin.RouterGroup) {
	router.GET("/register", h.login)
}

func (h *Handler) login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
