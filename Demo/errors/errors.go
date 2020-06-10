package errors

import "errors"

var (
	ErrUserWithUsernameAlreadyExist = errors.New("user with username already exists")
	ErrUserWithEmailAlreadyExists = errors.New("user with email already exists")
)

type MessageError struct {
	Message	string `json:"message"`
}
