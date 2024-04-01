package types

type AuthStore interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password, hashPassword string) error
	CreateJWT(secret []byte, userID int) (string, error)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" binding:"required`
	Password string `json:"password" form:"password" binding:"required"`
}
