package interfaceinfologic

import (
	"context"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInterfaceInfoLogic {
	return &AddInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddInterfaceInfoLogic) AddInterfaceInfo(in *pb.AddInterfaceInfoReq) (*pb.AddInterfaceInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddInterfaceInfoResp{}, nil
}
