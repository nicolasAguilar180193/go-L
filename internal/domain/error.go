package domain

import (
	"errors"
	"fmt"
)

const (
	ErrCodeDuplicateKey = "ERR_DUPLICATE_KEY"
	ErrCodeNotFound     = "ERR_NOT_FOUND"
)

var ErrDuplicateKey = errors.New("duplicate key error")

type AppError struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}
