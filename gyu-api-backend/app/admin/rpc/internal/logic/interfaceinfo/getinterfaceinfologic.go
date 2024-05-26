package interfaceinfologic

import (
	"context"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInterfaceInfoLogic {
	return &GetInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInterfaceInfoLogic) GetInterfaceInfo(in *pb.GetInterfaceInfoReq) (*pb.GetInterfaceInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetInterfaceInfoResp{}, nil
}
