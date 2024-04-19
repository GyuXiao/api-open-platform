package interfaceInfo

import (
	"context"
	"encoding/json"
	"github.com/GyuXiao/gyu-api-sdk/sdk"
	"github.com/GyuXiao/gyu-api-sdk/sdk/request"
	"github.com/GyuXiao/gyu-api-sdk/sdk/response"
	"github.com/GyuXiao/gyu-api-sdk/service/user"
	"github.com/zeromicro/go-zero/core/logc"
	"gyu-api-backend/app/admin/api/internal/models"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"strconv"
	"strings"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	accessKey = "gyu"
	secretKey = "abcdefg"
)

type OnlineInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnlineInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineInterfaceInfoLogic {
	return &OnlineInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnlineInterfaceInfoLogic) OnlineInterfaceInfo(req *types.OnlineInterfaceInfoReq) (resp *types.OnlineInterfaceInfoResp, err error) {
	// 0 通过 token 获取 redis 存储的 userRole，如果不是管理者，则不能执行上线操作
	token := strings.Split(req.Authorization, " ")[1]
	tokenLogic := models.NewDefaultTokenModel(l.svcCtx.RedisClient)
	result, err := tokenLogic.CheckTokenExist(token)
	if err != nil {
		return nil, err
	}
	userRoleStr := result[1]
	userRole, _ := strconv.Atoi(userRoleStr)
	if userRole != constant.AdminRole {
		return nil, xerr.NewErrCode(xerr.PermissionDenied)
	}
	// 1 校验接口是否存在（通过 id 查找接口）
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	_, err = interfaceInfoModel.SearchInterfaceInfoById(req.Id)
	if err != nil {
		return nil, err
	}
	// 2 校验接口是否可以调用
	// 具体做法是，调用 SDK 创建 client 并向目标 url 发起请求
	config := sdk.NewConfig(accessKey, secretKey)
	client, err := user.NewClient(config)
	if err != nil {
		logc.Infof(l.ctx, "SDK 创建客户端错误: %v", err)
		return nil, xerr.NewErrCode(xerr.SDKNewClientError)
	}
	userTest := user.NewUser("userTest1")
	userJson, _ := json.Marshal(userTest)
	baseReq := &request.BaseRequest{
		URL:    "http://127.0.0.1:8123/api/user",
		Method: "POST",
		Header: nil,
		Body:   string(userJson),
	}
	baseRsp := &response.BaseResponse{}
	err = client.Send(baseReq, baseRsp)
	if err != nil {
		logc.Infof(l.ctx, "向模拟接口发起请求错误: %v", err)
		return nil, xerr.NewErrCode(xerr.SDKSendRequestError)
	}
	// 3 修改接口状态为 online
	err = interfaceInfoModel.UpdateInterfaceInfoStatus(constant.Online, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.OnlineInterfaceInfoResp{IsOnline: true}, nil
}