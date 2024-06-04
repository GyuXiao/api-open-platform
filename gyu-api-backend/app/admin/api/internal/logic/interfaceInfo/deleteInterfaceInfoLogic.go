package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/common/userTools"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInterfaceInfoLogic {
	return &DeleteInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteInterfaceInfoLogic) DeleteInterfaceInfo(req *types.DeleteInterfaceInfoReq) (resp *types.DeleteInterfaceInfoResp, err error) {
	// 1,通过 token 获取 redis 存储的 userRole，如果不是管理者，则不能执行删除操作
	token := strings.Split(req.Authorization, " ")[1]
	err = userTools.CheckUserIsAdminRole(l.svcCtx.RedisClient, token)
	if err != nil {
		return nil, err
	}

	// 2,调用 rpc 模块的接口删除方法
	deleteInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.DeleteInterfaceInfo(l.ctx, &interfaceinfo.DeleteInterfaceInfoReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DeleteInterfaceInfoResp{IsDeleted: deleteInterfaceInfoResp.IsDeleted}, nil
}
