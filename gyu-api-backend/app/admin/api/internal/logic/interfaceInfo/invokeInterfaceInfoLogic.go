package interfaceInfo

import (
	"context"
	"github.com/GyuXiao/gyu-api-sdk/sdk"
	"github.com/GyuXiao/gyu-api-sdk/sdk/request"
	"github.com/GyuXiao/gyu-api-sdk/sdk/response"
	"github.com/zeromicro/go-zero/core/logc"
	"gyu-api-backend/app/admin/api/internal/models"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"strings"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	sdkService "github.com/GyuXiao/gyu-api-sdk/service/user"
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
	// 1 根据 id 查询接口：1，是否存在；2，是否已上线；
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	interfaceInfo, err := interfaceInfoModel.SearchInterfaceInfoById(req.Id)
	if err != nil {
		return nil, err
	}
	if interfaceInfo != nil && interfaceInfo.Status == constant.Offline {
		return nil, xerr.NewErrCode(xerr.InterfaceInfoOfflineError)
	}
	// 2 获取当前用户，通过 token 拿到 username，通过 username 查数据库拿到 ak 和 sk
	token := strings.Split(req.Authorization, " ")[1]
	tokenLogic := models.NewDefaultTokenModel(l.svcCtx.RedisClient)
	result, err := tokenLogic.CheckTokenExist(token)
	if err != nil {
		return nil, err
	}
	username := result[2]
	userModel := models.NewDefaultUserModel(l.svcCtx.DBEngin)
	user, err := userModel.SearchUserByUsername(username)
	if err != nil {
		return nil, err
	}
	// 3 再通过 ak 和 sk，调用 sdk 发起请求
	config := sdk.NewConfig(user.AccessKey, user.SecretKey)
	client, err := sdkService.NewClient(config)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.SDKNewClientError)
	}
	baseReq := &request.BaseRequest{
		URL:    constant.GatewayHost + "/api/invoke",
		Method: "POST",
		Header: nil,
		Body:   req.RequestParams,
	}
	baseRsp := &response.BaseResponse{}
	err = client.Send(baseReq, baseRsp)
	if err != nil {
		logc.Infof(l.ctx, "向模拟接口发起请求错误: %v", err)
		return nil, xerr.NewErrCode(xerr.SDKSendRequestError)
	}
	return &types.InvokeInterfaceInfoResp{ResponseObject: baseRsp}, nil
}
