package auth

import (
	"fmt"
	"net/http"

	"github.com/Kei-K23/go-otp/internal/types"
	"github.com/Kei-K23/go-otp/internal/utils"
	"github.com/Kei-K23/go-otp/templates/register"
	"github.com/Kei-K23/go-otp/templates/verify"
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
	router.GET("/verify", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", verify.Verify())
	})
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", register.Register())
	})
}

func (h *Handler) register(c *gin.Context) {
	var payload types.CreateUser
	if err := c.ShouldBind(&payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	randomUUID := uuid.New()
	payload.Token = randomUUID.String()[:8]

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

	// NOTE: This is a workaround for JSON responses
	// utils.WriteJSON(c, http.StatusCreated, gin.H{
	// 	"id":         user.ID,
	// 	"name":       user.Name,
	// 	"email":      user.Email,
	// 	"phone":      user.Phone,
	// 	"IsVerified": false,
	// 	"created_at": user.CreatedAt,
	// })

	c.Redirect(303, fmt.Sprintf("/api/v1/verify?userId=%d", user.ID))
}
