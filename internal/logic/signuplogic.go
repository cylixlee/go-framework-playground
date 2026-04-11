// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cylixlee/go-framework-playground/internal/model"
	"github.com/cylixlee/go-framework-playground/internal/svc"
	"github.com/cylixlee/go-framework-playground/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.SignUpRequest) (*types.SignUpResponse, error) {
	user := model.User{
		Username: sql.NullString{String: req.Username, Valid: true},
		Password: sql.NullString{String: req.Password, Valid: true},
	}
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &user)
	if err != nil {
		return nil, err
	}
	return &types.SignUpResponse{Message: fmt.Sprintf("Hello %s", req.Username)}, nil
}
