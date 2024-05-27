package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"strings"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnlineInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnlineInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineInterfaceInfoLogic {
	return &OnlineInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnlineInterfaceInfoLogic) OnlineInterfaceInfo(req *types.OnlineInterfaceInfoReq) (resp *types.OnlineInterfaceInfoResp, err error) {
	token := strings.Split(req.Authorization, " ")[1]
	onlineInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.OnlineInterfaceInfo(l.ctx, &interfaceinfo.OnlineInterfaceInfoReq{
		Id:        req.Id,
		AuthToken: token,
	})
	if err != nil {
		return nil, err
	}

	return &types.OnlineInterfaceInfoResp{IsOnline: onlineInterfaceInfoResp.IsOnline}, nil
}
