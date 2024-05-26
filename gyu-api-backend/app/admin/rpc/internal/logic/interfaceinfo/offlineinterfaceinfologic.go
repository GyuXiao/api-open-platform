package interfaceinfologic

import (
	"context"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OfflineInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOfflineInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OfflineInterfaceInfoLogic {
	return &OfflineInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OfflineInterfaceInfoLogic) OfflineInterfaceInfo(in *pb.OfflineInterfaceInfoReq) (*pb.OfflineInterfaceInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.OfflineInterfaceInfoResp{}, nil
}
