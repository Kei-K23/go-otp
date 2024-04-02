package todo

import (
	"net/http"

	"github.com/Kei-K23/go-otp/internal/middlewares"
	"github.com/Kei-K23/go-otp/internal/types"
	"github.com/Kei-K23/go-otp/internal/utils"
	"github.com/Kei-K23/go-otp/templates/components/todo_comp"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	todoStore types.TodoStore
}

func NewHandler(todoStore types.TodoStore) *Handler {
	return &Handler{todoStore: todoStore}
}

func (h *Handler) RegisterRoutes(router gin.RouterGroup) {
	router.POST("/todos", h.createNewTodo)
}

func (h *Handler) createNewTodo(c *gin.Context) {

	var payload types.CreateTodo
	if err := c.ShouldBind(&payload); err != nil {
		utils.WriteError(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	_, err := h.todoStore.CreateNewTodo(payload, id)

	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todos, err := h.todoStore.GetAllTodoByUserId(id)

	if err != nil {
		utils.WriteError(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "", todo_comp.TodoComp(todos))
}
