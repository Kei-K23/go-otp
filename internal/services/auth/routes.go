package auth

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kei-K23/go-otp/internal/types"
	"github.com/Kei-K23/go-otp/internal/utils"
	"github.com/Kei-K23/go-otp/templates/login"
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
	router.POST("/verify", h.verify)
	router.POST("/login", h.login)

	// template rendering here
	router.GET("/verify", func(c *gin.Context) {

		id := c.Query("userId")
		statusErr := c.Query("error")

		uID, err := strconv.Atoi(id)

		if err != nil {
			utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "", verify.Verify(fmt.Sprintf("/api/v1/verify?userId=%d", uID), statusErr))
	})

	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", register.Register())
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", login.Login())
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

func (h *Handler) verify(c *gin.Context) {
	var payload types.VerifyUser
	id := c.Query("userId")
	fmt.Println(id)
	uID, err := strconv.Atoi(id)

	if err != nil {
		c.Redirect(303, fmt.Sprintf("/api/v1/verify?userId=%d&error=error", uID))
		return
	}

	if err := c.ShouldBind(&payload); err != nil {
		c.Redirect(303, fmt.Sprintf("/api/v1/verify?userId=%d&error=error", uID))
		return
	}

	err = h.userStore.VerifyUserAcc(uID, payload.Token)

	if err != nil {
		fmt.Println(err)
		c.Redirect(303, fmt.Sprintf("/api/v1/verify?userId=%d&error=error", uID))
		return
	}

	c.Redirect(303, "/api/v1/login")
}

func (h *Handler) login(c *gin.Context) {
	var payload types.VerifyUser
	id := c.Query("userId")

	uID, err := strconv.Atoi(id)

	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBind(&payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.userStore.VerifyUserAcc(uID, payload.Token)

	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(303, "/api/v1/login")
}
