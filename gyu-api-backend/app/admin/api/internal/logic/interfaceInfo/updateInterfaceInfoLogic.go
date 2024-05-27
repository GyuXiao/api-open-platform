package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInterfaceInfoLogic {
	return &UpdateInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInterfaceInfoLogic) UpdateInterfaceInfo(req *types.UpdateInterfaceInfoReq) (resp *types.UpdateInterfaceInfoResp, err error) {
	// todo：补充校验参数逻辑
	// todo: 管理员才能更新接口信息

	updateInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.UpdateInterfaceInfo(l.ctx, &interfaceinfo.UpdateInterfaceInfoReq{
		Name:           req.Name,
		Description:    req.Description,
		Url:            req.Url,
		RequestHeader:  req.RequestHeader,
		ResponseHeader: req.ResponseHeader,
		Method:         req.Method,
		RequestParams:  req.RequestParams,
		Id:             req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdateInterfaceInfoResp{IsUpdated: updateInterfaceInfoResp.IsUpdated}, nil
}
