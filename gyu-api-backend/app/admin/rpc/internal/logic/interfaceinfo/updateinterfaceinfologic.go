package interfaceinfologic

import (
	"context"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInterfaceInfoLogic {
	return &UpdateInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateInterfaceInfoLogic) UpdateInterfaceInfo(in *pb.UpdateInterfaceInfoReq) (*pb.UpdateInterfaceInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateInterfaceInfoResp{}, nil
}
