package greetlogic

import (
	"context"

	"github.com/cylixlee/go-framework-playground/internal/proto"
	"github.com/cylixlee/go-framework-playground/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *proto.GreetRequest) (*proto.GreetResponse, error) {
	// todo: add your logic here and delete this line

	return &proto.GreetResponse{}, nil
}
