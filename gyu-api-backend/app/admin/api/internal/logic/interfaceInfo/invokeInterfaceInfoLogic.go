package interfaceInfo

import (
	"context"
	"github.com/GyuXiao/gyu-api-sdk/sdk/response"
	"github.com/duke-git/lancet/v2/maputil"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/common/tools"
	"strings"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvokeInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInvokeInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvokeInterfaceInfoLogic {
	return &InvokeInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InvokeInterfaceInfoLogic) InvokeInterfaceInfo(req *types.InvokeInterfaceInfoReq) (resp *types.InvokeInterfaceInfoResp, err error) {
	token := strings.Split(req.Authorization, " ")[1]
	invokeInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.InvokeInterfaceInfo(l.ctx, &interfaceinfo.InvokeInterfaceInfoReq{
		Id:            req.Id,
		RequestParams: req.RequestParams,
		AuthToken:     token,
	})
	if err != nil {
		return nil, err
	}

	var baseRsp response.BaseResponse
	mp := tools.MapConvertStringToAny(invokeInterfaceInfoResp.ResponseObject)
	_ = maputil.MapToStruct(mp, &baseRsp)

	return &types.InvokeInterfaceInfoResp{ResponseObject: baseRsp}, nil
}
