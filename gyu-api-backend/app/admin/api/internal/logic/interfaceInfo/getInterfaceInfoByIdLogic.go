package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInterfaceInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInterfaceInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInterfaceInfoByIdLogic {
	return &GetInterfaceInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInterfaceInfoByIdLogic) GetInterfaceInfoById(req *types.GetInterfaceInfoReq) (resp *types.GetInterfaceInfoResp, err error) {
	getInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.GetInterfaceInfo(l.ctx, &interfaceinfo.GetInterfaceInfoReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.GetInterfaceInfoResp{
		Description:    getInterfaceInfoResp.Description,
		Url:            getInterfaceInfoResp.Url,
		RequestParams:  getInterfaceInfoResp.RequestParams,
		RequestHeader:  getInterfaceInfoResp.RequestHeader,
		ResponseHeader: getInterfaceInfoResp.ResponseHeader,
		Status:         uint8(getInterfaceInfoResp.Status),
		Method:         getInterfaceInfoResp.Method,
		CreateTime:     getInterfaceInfoResp.CreateTime,
		UpdateTime:     getInterfaceInfoResp.UpdateTime,
	}, nil
}
