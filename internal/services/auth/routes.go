package auth

import (
	"net/http"

	"github.com/Kei-K23/go-otp/internal/types"
	"github.com/Kei-K23/go-otp/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	authStore types.AuthStore
	userStore types.UserStore
}

type Login struct {
	Username string `json:"username" binding:"required"`
}

func NewHandler(authStore types.AuthStore, userStore types.UserStore) *Handler {
	return &Handler{authStore: authStore, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router gin.RouterGroup) {
	router.POST("/register", h.register)
}

func (h *Handler) register(c *gin.Context) {
	var payload types.CreateUser
	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payload.Token = uuid.NewString()
	hashPassword, err := h.authStore.HashPassword(payload.Password)

	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	payload.Password = hashPassword

	user, err := h.userStore.CreateUser(payload)

	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.WriteJSON(c, http.StatusCreated, gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"phone":      user.Phone,
		"IsVerified": false,
		"created_at": user.CreatedAt,
	})
}
