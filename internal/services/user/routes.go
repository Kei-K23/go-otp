package user

import (
	"github.com/Kei-K23/go-otp/internal/types"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userStore types.UserStore
}

type Login struct {
	Username string `json:"username" binding:"required"`
}

func NewHandler(userStore types.UserStore) *Handler {
	return &Handler{userStore: userStore}
}

func (h *Handler) RegisterRoutes(router gin.RouterGroup) {
	// router.GET("/register", h.login)
}
