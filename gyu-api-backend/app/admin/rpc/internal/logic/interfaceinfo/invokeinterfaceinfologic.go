package interfaceinfologic

import (
	"context"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvokeInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvokeInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvokeInterfaceInfoLogic {
	return &InvokeInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvokeInterfaceInfoLogic) InvokeInterfaceInfo(in *pb.InvokeInterfaceInfoReq) (*pb.InvokeInterfaceInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.InvokeInterfaceInfoResp{}, nil
}
