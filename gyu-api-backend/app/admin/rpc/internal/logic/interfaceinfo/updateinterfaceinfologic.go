package interfaceinfologic

import (
	"context"
	"gyu-api-backend/app/admin/models"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInterfaceInfoLogic {
	return &UpdateInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateInterfaceInfoLogic) UpdateInterfaceInfo(in *pb.UpdateInterfaceInfoReq) (*pb.UpdateInterfaceInfoResp, error) {
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	interfaceMap := map[string]interface{}{
		"name":           in.Name,
		"description":    in.Description,
		"url":            in.Url,
		"requestHeader":  in.RequestHeader,
		"responseHeader": in.ResponseHeader,
		"method":         in.Method,
		"requestParams":  in.RequestParams,
	}
	err := interfaceInfoModel.UpdateInterfaceInfo(in.Id, interfaceMap)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateInterfaceInfoResp{IsUpdated: true}, nil
}
