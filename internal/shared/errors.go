package shared

import "errors"

var (
	ErrAlreadyExists = errors.New("record already exists")
	ErrNotFound      = errors.New("record not found")
	ErrBadInput      = errors.New("bad input")
)
