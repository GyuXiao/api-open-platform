package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"strings"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OfflineInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOfflineInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OfflineInterfaceInfoLogic {
	return &OfflineInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OfflineInterfaceInfoLogic) OfflineInterfaceInfo(req *types.OfflineInterfaceInfoReq) (resp *types.OfflineInterfaceInfoResp, err error) {
	token := strings.Split(req.Authorization, " ")[1]
	offlineInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.OfflineInterfaceInfo(l.ctx, &interfaceinfo.OfflineInterfaceInfoReq{
		Id:        req.Id,
		AuthToken: token,
	})
	if err != nil {
		return nil, err
	}

	return &types.OfflineInterfaceInfoResp{IsOffline: offlineInterfaceInfoResp.IsOffline}, nil
}
