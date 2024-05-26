package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"strconv"
	"strings"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OfflineInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOfflineInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OfflineInterfaceInfoLogic {
	return &OfflineInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OfflineInterfaceInfoLogic) OfflineInterfaceInfo(req *types.OfflineInterfaceInfoReq) (resp *types.OfflineInterfaceInfoResp, err error) {
	// 0 通过 token 获取 redis 存储的 userRole，如果不是管理者，则不能执行下线操作
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
	// 2 修改接口状态为 offline
	err = interfaceInfoModel.UpdateInterfaceInfoStatus(constant.Offline, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.OfflineInterfaceInfoResp{IsOffline: true}, nil
}
