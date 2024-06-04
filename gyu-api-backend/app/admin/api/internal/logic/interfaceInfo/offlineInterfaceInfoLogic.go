package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/common/userTools"
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
	// 0 通过 token 获取 redis 存储的 userRole，如果不是管理者，则不能执行下线操作
	token := strings.Split(req.Authorization, " ")[1]
	err = userTools.CheckUserIsAdminRole(l.svcCtx.RedisClient, token)
	if err != nil {
		return nil, err
	}

	offlineInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.OfflineInterfaceInfo(l.ctx, &interfaceinfo.OfflineInterfaceInfoReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.OfflineInterfaceInfoResp{IsOffline: offlineInterfaceInfoResp.IsOffline}, nil
}
