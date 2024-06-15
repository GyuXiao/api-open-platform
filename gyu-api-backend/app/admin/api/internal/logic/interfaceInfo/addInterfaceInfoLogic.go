package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInterfaceInfoLogic {
	return &AddInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddInterfaceInfoLogic) AddInterfaceInfo(req *types.AddInterfaceInfoReq) (resp *types.AddInterfaceInfoResp, err error) {
	// 1,校验参数逻辑
	if req.UserId == constant.BlankInt || req.Name == constant.BlankString {
		return nil, xerr.NewErrCode(xerr.RequestParamError)
	}

	// 2,调用 rpc 模块的新增接口方法
	addInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.AddInterfaceInfo(l.ctx, &interfaceinfo.AddInterfaceInfoReq{
		Name:           req.Name,
		Description:    req.Description,
		Url:            req.Url,
		RequestParams:  req.RequestParams,
		RequestHeader:  req.RequestHeader,
		ResponseHeader: req.ResponseHeader,
		Method:         req.Method,
		UserId:         req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddInterfaceInfoResp{Name: addInterfaceInfoResp.Name}, nil
}
