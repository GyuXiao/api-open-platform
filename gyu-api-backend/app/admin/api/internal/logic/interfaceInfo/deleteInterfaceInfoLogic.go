package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"

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

	deleteInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.DeleteInterfaceInfo(l.ctx, &interfaceinfo.DeleteInterfaceInfoReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.DeleteInterfaceInfoResp{IsDeleted: deleteInterfaceInfoResp.IsDeleted}, nil
}
