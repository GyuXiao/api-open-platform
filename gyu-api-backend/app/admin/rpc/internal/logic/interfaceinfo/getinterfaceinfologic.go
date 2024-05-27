package interfaceinfologic

import (
	"context"
	"gyu-api-backend/app/admin/models"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInterfaceInfoLogic {
	return &GetInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInterfaceInfoLogic) GetInterfaceInfo(in *pb.GetInterfaceInfoReq) (*pb.GetInterfaceInfoResp, error) {
	interfaceInfoLogic := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	interfaceInfo, err := interfaceInfoLogic.SearchInterfaceInfoById(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetInterfaceInfoResp{
		Description:    interfaceInfo.Description,
		Url:            interfaceInfo.Url,
		RequestParams:  interfaceInfo.RequestParams,
		RequestHeader:  interfaceInfo.RequestHeader,
		ResponseHeader: interfaceInfo.ResponseHeader,
		Status:         uint32(interfaceInfo.Status),
		Method:         interfaceInfo.Method,
		CreateTime:     interfaceInfo.CreateTime,
		UpdateTime:     interfaceInfo.UpdateTime,
	}, nil
}
