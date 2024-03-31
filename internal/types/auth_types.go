package types

type AuthStore interface {
	HashPassword(password string) (string, error)
}
