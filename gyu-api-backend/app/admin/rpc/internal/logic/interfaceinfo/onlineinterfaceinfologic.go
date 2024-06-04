package interfaceinfologic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"
	"gyu-api-backend/common/constant"
)

type OnlineInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOnlineInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineInterfaceInfoLogic {
	return &OnlineInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OnlineInterfaceInfoLogic) OnlineInterfaceInfo(in *pb.OnlineInterfaceInfoReq) (*pb.OnlineInterfaceInfoResp, error) {
	// 1,校验接口是否存在（通过 id 查找接口）
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	_, err := interfaceInfoModel.SearchInterfaceInfoById(in.Id)
	if err != nil {
		return nil, err
	}

	// 2,修改接口状态为 online
	err = interfaceInfoModel.UpdateInterfaceInfoStatus(constant.Online, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.OnlineInterfaceInfoResp{IsOnline: true}, nil
}
