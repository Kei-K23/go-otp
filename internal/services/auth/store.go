package auth

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Store struct {
	db *sql.DB
}

type JWTClaim struct {
	UserID  int
	Expires int64
	jwt.RegisteredClaims
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashPassword), err
}

func (s *Store) VerifyPassword(password, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) CreateJWT(secret []byte, userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaim{
		UserID:  userID,
		Expires: int64(time.Second * time.Duration(3600*24*7)),
	})

	t, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return t, nil
}
