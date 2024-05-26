package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/models"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInterfaceInfoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInterfaceInfoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInterfaceInfoByIdLogic {
	return &GetInterfaceInfoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetInterfaceInfoByIdLogic) GetInterfaceInfoById(req *types.GetInterfaceInfoReq) (resp *types.GetInterfaceInfoResp, err error) {
	interfaceInfoLogic := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	interfaceInfo, err := interfaceInfoLogic.SearchInterfaceInfoById(req.Id)
	if err != nil {
		return nil, err
	}
	return &types.GetInterfaceInfoResp{
		Description:    interfaceInfo.Description,
		Url:            interfaceInfo.Url,
		RequestParams:  interfaceInfo.RequestParams,
		RequestHeader:  interfaceInfo.RequestHeader,
		ResponseHeader: interfaceInfo.ResponseHeader,
		Status:         interfaceInfo.Status,
		Method:         interfaceInfo.Method,
		CreateTime:     interfaceInfo.CreateTime,
		UpdateTime:     interfaceInfo.UpdateTime,
	}, nil
}
