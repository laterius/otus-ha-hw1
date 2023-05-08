package domain

import "github.com/pkg/errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUserId     = errors.New("invalid user id")
	_                    = errors.New("internal error")
	ErrInvalidPassHash   = errors.New("invalid password hash")
)
