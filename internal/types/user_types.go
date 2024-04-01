package types

type UserStore interface {
	GetUserById(id int64) (*User, error)
	CreateUser(cU CreateUser) (*User, error)
	VerifyUserAcc(uID int, token string) error
	GetUserByEmail(email string) (*User, error)
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Token      string `json:"token"`
	IsVerified bool   `json:is_verified"`
	CreatedAt  string `json:"created_at"`
}

type CreateUser struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required`
	Password string `json:"password" form:"password" binding:"required"`
	Phone    string `json:"phone" form:"phone" binding:"required`
	Token    string `json:"token"`
}

type VerifyUser struct {
	Token string `json:"token" form:"token" binding:"required"`
}

type UpdateUser struct {
	Name       string `json:"name" binding:"optional"`
	Email      string `json:"email" binding:"optional"`
	Password   string `json:"password" binding:"optional`
	Phone      string `json:"phone" binding:"optional"`
	Token      string `json:"token" binding:"optional"`
	IsVerified bool   `json:is_verified" binding:"optional"`
}
