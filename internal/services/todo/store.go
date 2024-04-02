package todo

import (
	"database/sql"

	"github.com/Kei-K23/go-otp/internal/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetTodoById(id int) (*types.Todo, error) {
	var t types.Todo
	stmt, err := s.db.Prepare("SELECT * FROM todos WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&t.ID, &t.Todo, &t.Completed, &t.UserID, &t.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
func (s *Store) GetAllTodoByUserId(userId int) ([]types.Todo, error) {
	var todos []types.Todo

	stmt, err := s.db.Prepare("SELECT * FROM todos WHERE user_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t types.Todo
		if err := rows.Scan(&t.ID, &t.Todo, &t.Completed, &t.UserID, &t.CreatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *Store) CreateNewTodo(todo types.CreateTodo, userID int) (*types.Todo, error) {

	stmt, err := s.db.Prepare("INSERT INTO todos (todo, user_id) VALUES (?, ?)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(todo.Todo, userID)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	t, err := s.GetTodoById(int(id))

	if err != nil {
		return nil, err
	}

	return t, nil

}
