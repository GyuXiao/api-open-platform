package interfaceinfologic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/common/xerr"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInterfaceInfoLogic {
	return &AddInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddInterfaceInfoLogic) AddInterfaceInfo(in *pb.AddInterfaceInfoReq) (*pb.AddInterfaceInfoResp, error) {
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	record, _ := interfaceInfoModel.SearchInterfaceInfoByName(in.Name)
	if record != nil {
		logc.Info(l.ctx, "接口名称已经存在，添加失败")
		return nil, xerr.NewErrCode(xerr.RecordDuplicateError)
	}
	interfaceMap := map[string]interface{}{
		"name":           in.Name,
		"description":    in.Description,
		"url":            in.Url,
		"requestParams":  in.RequestParams,
		"requestHeader":  in.RequestHeader,
		"responseHeader": in.ResponseHeader,
		"method":         in.Method,
		"status":         0,
		"userId":         in.UserId,
	}
	err := interfaceInfoModel.AddInterfaceInfo(interfaceMap)
	if err != nil {
		return nil, err
	}
	return &pb.AddInterfaceInfoResp{Name: in.Name}, nil
}
