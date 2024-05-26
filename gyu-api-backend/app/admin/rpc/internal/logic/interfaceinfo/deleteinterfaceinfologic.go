package interfaceinfologic

import (
	"context"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInterfaceInfoLogic {
	return &DeleteInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteInterfaceInfoLogic) DeleteInterfaceInfo(in *pb.DeleteInterfaceInfoReq) (*pb.DeleteInterfaceInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DeleteInterfaceInfoResp{}, nil
}
