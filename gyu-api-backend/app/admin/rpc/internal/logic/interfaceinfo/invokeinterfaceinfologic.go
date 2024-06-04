package interfaceinfologic

import (
	"context"
	"github.com/GyuXiao/gyu-api-sdk/client"
	"github.com/GyuXiao/gyu-api-sdk/sdk"
	"github.com/GyuXiao/gyu-api-sdk/sdk/request"
	"github.com/GyuXiao/gyu-api-sdk/sdk/response"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/tools"
	"gyu-api-backend/common/userTools"
	"gyu-api-backend/common/xerr"
	"strconv"
)

type InvokeInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvokeInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvokeInterfaceInfoLogic {
	return &InvokeInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvokeInterfaceInfoLogic) InvokeInterfaceInfo(in *pb.InvokeInterfaceInfoReq) (*pb.InvokeInterfaceInfoResp, error) {
	// 1 根据 id 查询接口：1，是否存在；2，是否已上线；
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	interfaceInfo, err := interfaceInfoModel.SearchInterfaceInfoById(in.Id)
	if err != nil {
		return nil, err
	}
	if interfaceInfo != nil && interfaceInfo.Status == constant.Offline {
		return nil, xerr.NewErrCode(xerr.InterfaceInfoOfflineError)
	}

	// 2 获取当前用户，通过 token 拿到 username，通过 username 查数据库拿到 ak 和 sk
	userInfo, err := userTools.GetUserInfo(l.svcCtx.RedisClient, in.AuthToken)
	if err != nil {
		return nil, err
	}
	username := userInfo[constant.KeyUsername]
	userModel := models.NewDefaultUserModel(l.svcCtx.DBEngin)
	user, err := userModel.SearchUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// 3 再通过 ak 和 sk，调用 sdk 发起请求
	config := sdk.NewConfig(user.AccessKey, user.SecretKey)
	clt, err := client.NewClient(config)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.SDKNewClientError)
	}
	itfId := strconv.Itoa(int(interfaceInfo.Id))
	baseReq := &request.BaseRequest{
		URL:    constant.GatewayHost + constant.GatewayUrl,
		Method: constant.PostMethod,
		Header: nil,
		ItfId:  itfId,
		Body:   in.RequestParams,
	}
	baseRsp := &response.BaseResponse{}
	err = clt.Send(baseReq, baseRsp)
	if err != nil {
		logc.Infof(l.ctx, "向模拟接口发起请求错误: %v", err)
		return nil, xerr.NewErrCode(xerr.SDKSendRequestError)
	}

	// 将 baseRsp.ErrorResponse 转换成 map[string]string
	resultMp := tools.StructConvertMap(baseRsp)

	return &pb.InvokeInterfaceInfoResp{ResponseObject: resultMp}, nil
}
