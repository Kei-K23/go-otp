package user

import (
	"net/http"

	"github.com/Kei-K23/go-otp/internal/middlewares"
	"github.com/Kei-K23/go-otp/internal/types"
	"github.com/Kei-K23/go-otp/internal/utils"
	"github.com/Kei-K23/go-otp/templates/users_template"
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
	router.GET("/users", func(c *gin.Context) {

		// Retrieve user ID from the context
		userID, exists := c.Get(string(middlewares.ClaimsContextKey))
		if !exists {
			c.Redirect(303, "/api/v1/login")
			return
		}

		// Convert userID to the appropriate type
		id, ok := userID.(int)
		if !ok {
			utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": "user ID is not of type int64"})
			return
		}

		user, err := h.userStore.GetUserById(int64(id))

		if err != nil {
			utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if !user.IsVerified {
			utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": "this user account is not verify yet! Please verify first to continue."})
			return
		}

		c.HTML(http.StatusOK, "", users_template.Users(*user))
	})
}
