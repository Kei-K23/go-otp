package user

import (
	"net/http"
	"strconv"

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
	router.GET("/users/:userId", func(c *gin.Context) {

		userId := c.Param("userId")

		id, err := strconv.Atoi(userId)

		if err != nil {
			utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": err})
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
