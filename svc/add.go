package svc

import (
	"context"
	"errors"
)

var ErrEmptyString = errors.New("cannot concat two empty strings")

type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}
