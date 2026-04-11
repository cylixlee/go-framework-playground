// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package logic

import (
	"context"

	"github.com/cylixlee/go-framework-playground/internal/svc"
	"github.com/cylixlee/go-framework-playground/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInRequest) (resp *types.SignInResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
