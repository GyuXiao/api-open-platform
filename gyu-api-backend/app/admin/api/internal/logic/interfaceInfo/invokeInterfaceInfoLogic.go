package interfaceInfo

import (
	"context"
	"github.com/GyuXiao/gyu-api-sdk/sdk/response"
	"github.com/zeromicro/go-zero/core/logc"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/tools"
	"gyu-api-backend/common/xerr"
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
	// 参数校验
	if req.Id <= constant.BlankInt {
		logc.Info(l.ctx, "调用接口时 id 不能为非正数")
		return nil, xerr.NewErrCode(xerr.RequestParamError)
	}
	token := strings.Split(req.Authorization, " ")[1]
	if token == constant.BlankString {
		logc.Info(l.ctx, "调用接口时 auth token 不能为空")
		return nil, xerr.NewErrCode(xerr.RequestParamError)
	}

	invokeInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.InvokeInterfaceInfo(l.ctx, &interfaceinfo.InvokeInterfaceInfoReq{
		Id:            req.Id,
		RequestParams: req.RequestParams,
		AuthToken:     token,
	})
	if err != nil {
		return nil, err
	}

	// 获取响应对象
	respObj := invokeInterfaceInfoResp.ResponseObject
	// 将 map[string]string 转换为 struct 对象
	var baseRsp response.ErrorResponse
	tools.MapConvertStruct(respObj, &baseRsp)

	return &types.InvokeInterfaceInfoResp{ResponseObject: baseRsp}, nil
}
