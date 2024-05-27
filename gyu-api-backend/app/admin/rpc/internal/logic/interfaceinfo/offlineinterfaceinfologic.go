package interfaceinfologic

import (
	"context"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"strconv"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type OfflineInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOfflineInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OfflineInterfaceInfoLogic {
	return &OfflineInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OfflineInterfaceInfoLogic) OfflineInterfaceInfo(in *pb.OfflineInterfaceInfoReq) (*pb.OfflineInterfaceInfoResp, error) {
	// 0 通过 token 获取 redis 存储的 userRole，如果不是管理者，则不能执行下线操作
	tokenLogic := models.NewDefaultTokenModel(l.svcCtx.RedisClient)
	result, err := tokenLogic.CheckTokenExist(in.AuthToken)
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
	_, err = interfaceInfoModel.SearchInterfaceInfoById(in.Id)
	if err != nil {
		return nil, err
	}

	// 2 修改接口状态为 offline
	err = interfaceInfoModel.UpdateInterfaceInfoStatus(constant.Offline, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.OfflineInterfaceInfoResp{IsOffline: true}, nil
}
