package shared

import (
	"errors"
	"fmt"
)

var (
	ErrValidation          = errors.New("validation error")
	ErrBadRequest          = errors.New("bad request")
	ErrRecordAlreadyExists = errors.New("record already exists")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrNotFound            = errors.New("not found")
)

func NewPgRepoErrQueryCreation(err error) error {
	return fmt.Errorf("failed to create the sql query: %w", err)
}

func NewPgRepoErrExecution(sql string, args any, err error) error {
	return fmt.Errorf("failed to execute the following query: %s - %v - %w", sql, args, err)
}

func NewPgRepoErrTransaction(stage string, err error) error {
	return fmt.Errorf("failed to %s the transaction: - %w", stage, err)
}

func NewPgRepoErrExistingRecord(name string, record string) error {
	return fmt.Errorf("%s '%s' - %w", name, record, ErrRecordAlreadyExists)
}
