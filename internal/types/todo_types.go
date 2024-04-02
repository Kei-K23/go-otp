package types

type TodoStore interface {
	CreateNewTodo(todo CreateTodo, userID int) (*Todo, error)
	GetTodoById(id int) (*Todo, error)
	GetAllTodoByUserId(userId int) ([]Todo, error)
}

type Todo struct {
	ID        int    `json:"id"`
	Todo      string `json:"title"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type CreateTodo struct {
	Todo string `json:"todo" form:"todo" binding:"required"`
}

type UpdateTodo struct {
	Todo string `json:"todo" form:"todo" binding:"required"`
}
