package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/userTools"
	"gyu-api-backend/common/xerr"
	"strings"

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
	// 1,校验参数逻辑
	if req.Id == constant.BlankInt || req.Name == constant.BlankString {
		return nil, xerr.NewErrCode(xerr.RequestParamError)
	}

	// 2,通过 token 获取 redis 存储的 userRole，如果不是管理者，则不能执行更新操作
	token := strings.Split(req.Authorization, " ")[1]
	err = userTools.CheckUserIsAdminRole(l.svcCtx.RedisClient, token)
	if err != nil {
		return nil, err
	}

	// 3,调用 rpc 模块的更新接口信息方法
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
