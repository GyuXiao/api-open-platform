package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/models"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInterfaceInfoLogic {
	return &UpdateInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInterfaceInfoLogic) UpdateInterfaceInfo(req *types.UpdateInterfaceInfoReq) (resp *types.UpdateInterfaceInfoResp, err error) {
	// todo：校验参数
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	interfaceMap := map[string]interface{}{
		"name":           req.Name,
		"description":    req.Description,
		"url":            req.Url,
		"requestHeader":  req.RequestHeader,
		"responseHeader": req.ResponseHeader,
		"method":         req.Method,
	}
	err = interfaceInfoModel.UpdateInterfaceInfo(req.Id, interfaceMap)
	if err != nil {
		return nil, err
	}
	return &types.UpdateInterfaceInfoResp{IsUpdated: true}, nil
}
