package interfaceInfo

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"gyu-api-backend/app/admin/api/internal/models"
	"gyu-api-backend/common/xerr"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddInterfaceInfoLogic {
	return &AddInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddInterfaceInfoLogic) AddInterfaceInfo(req *types.AddInterfaceInfoReq) (resp *types.AddInterfaceInfoResp, err error) {
	// todo1: 参数校验
	interfaceInfoModel := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	record, _ := interfaceInfoModel.SearchInterfaceInfoByName(req.Name)
	if record != nil {
		logc.Info(l.ctx, "接口名称已经存在，添加失败")
		return nil, xerr.NewErrCode(xerr.RecordDuplicateError)
	}
	interfaceMap := map[string]interface{}{
		"name":           req.Name,
		"description":    req.Description,
		"url":            req.Url,
		"requestParams":  req.RequestParams,
		"requestHeader":  req.RequestHeader,
		"responseHeader": req.ResponseHeader,
		"method":         req.Method,
		"status":         0,
		"userId":         req.UserId,
	}
	err = interfaceInfoModel.AddInterfaceInfo(interfaceMap)
	if err != nil {
		return nil, err
	}
	return &types.AddInterfaceInfoResp{Name: req.Name}, nil
}
