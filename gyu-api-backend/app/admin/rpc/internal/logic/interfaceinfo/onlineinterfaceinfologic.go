package interfaceinfologic

import (
	"context"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnlineInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOnlineInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineInterfaceInfoLogic {
	return &OnlineInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OnlineInterfaceInfoLogic) OnlineInterfaceInfo(in *pb.OnlineInterfaceInfoReq) (*pb.OnlineInterfaceInfoResp, error) {
	// todo: add your logic here and delete this line

	return &pb.OnlineInterfaceInfoResp{}, nil
}
