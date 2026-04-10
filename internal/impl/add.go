package impl

import (
	"context"

	"github.com/cylixlee/go-framework-playground/svc"
)

type AddServiceImpl struct{}

func (AddServiceImpl) Sum(ctx context.Context, a, b int) (int, error) {
	return a + b, nil
}

func (AddServiceImpl) Concat(ctx context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", svc.ErrEmptyString
	}
	return a + b, nil
}
