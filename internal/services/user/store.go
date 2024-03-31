package user

import (
	"database/sql"
	"fmt"

	"github.com/Kei-K23/go-otp/internal/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserById(id int64) (*types.User, error) {
	var user types.User

	stmt, err := s.db.Prepare("SELECT * FROM users WHERE id = ?")

	if err != nil {
		return nil, fmt.Errorf("error preparing user: %v", err)
	}

	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Token, &user.IsVerified, &user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("error executing user: %v", err)
	}

	return &user, nil
}

func (s *Store) CreateUser(cU types.CreateUser) (*types.User, error) {
	stmt, err := s.db.Prepare("INSERT INTO users (name, email, password, phone, token) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return nil, fmt.Errorf("error preparing user: %v", err)
	}

	res, err := stmt.Exec(cU.Name, cU.Email, cU.Password, cU.Phone, cU.Token)

	if err != nil {
		return nil, fmt.Errorf("error executing user: %v", err)
	}

	uID, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	user, err := s.GetUserById(uID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving created user: %v", err)
	}

	return user, nil
}
