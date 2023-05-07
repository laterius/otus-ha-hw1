package domain

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidUserId   = errors.New("invalid user id")
	_                  = errors.New("internal error")
	ErrInvalidPassHash = errors.New("invalid password hash")
)
