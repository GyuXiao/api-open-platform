package interfaceinfologic

import (
	"context"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"
	"gyu-api-backend/common/constant"

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

// 1，token 的身份校验应该放在 api 模块执行

func (l *OfflineInterfaceInfoLogic) OfflineInterfaceInfo(in *pb.OfflineInterfaceInfoReq) (*pb.OfflineInterfaceInfoResp, error) {
	// 1 校验接口是否存在（通过 id 查找接口）
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	_, err := interfaceInfoModel.SearchInterfaceInfoById(in.Id)
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
