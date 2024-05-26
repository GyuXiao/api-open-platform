package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/models"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteInterfaceInfoLogic {
	return &DeleteInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteInterfaceInfoLogic) DeleteInterfaceInfo(req *types.DeleteInterfaceInfoReq) (resp *types.DeleteInterfaceInfoResp, err error) {
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	err = interfaceInfoModel.DeleteInterfaceInfo(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.DeleteInterfaceInfoResp{IsDeleted: true}, nil
}
