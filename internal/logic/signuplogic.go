// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

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

func (l *SignUpLogic) SignUp(req *types.SignUpRequest) (resp *types.SignUpResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
