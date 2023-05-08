package domain

import "github.com/laterius/service_architecture_hw3/app/pkg/types"

type UserId int64
type Username string
type Password string

func (t UserId) Validate() error {
	if t <= 0 {
		return ErrInvalidUserId
	}
	return nil
}

type User struct {
	Id           UserId
	Username     string `db:"username"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	Email        string `db:"email"`
	Phone        string `db:"phone"`
	Age          int64  `db:"age"`
	Gender       string `db:"gender"`
	Hobby        string `db:"hobby"`
	City         string `db:"city"`
	Password     string
	PasswordHash string `db:"password_hash"`
	Remember     string `db:"remember"`
	RememberHash string `db:"remember_hash"`
}

type UserPartialData = types.Kv
