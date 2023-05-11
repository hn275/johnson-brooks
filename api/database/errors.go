package database

import (
	"errors"
)

var (
	ErrNoDocument = errors.New("document not found")
	ErrWriteFail  = errors.New("failed to write to database")
)
