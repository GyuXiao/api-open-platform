package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/models"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPageListLogic {
	return &GetPageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPageListLogic) GetPageList(req *types.PageListReq) (resp *types.PageListResp, err error) {
	interfaceInfoLogic := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	result, total, err := interfaceInfoLogic.FindListPage(req.Keyword, req.Current, req.PageSize)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	var interfaceInfoList []types.InterfaceInfo
	for _, interfaceInfo := range result {
		tmp := types.InterfaceInfo{}
		copier.Copy(&tmp, interfaceInfo)
		interfaceInfoList = append(interfaceInfoList, tmp)
	}
	return &types.PageListResp{
		Total:   uint64(total),
		Records: interfaceInfoList,
	}, nil
}
