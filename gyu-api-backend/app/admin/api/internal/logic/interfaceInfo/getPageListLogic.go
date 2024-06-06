package interfaceInfo

import (
	"context"
	"github.com/jinzhu/copier"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

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

	interfaceListResp, err := l.svcCtx.InterfaceInfoRpc.GetPageList(l.ctx, &interfaceinfo.PageListReq{
		Keyword:  req.Name,
		Current:  req.Current,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var list []types.InterfaceInfo
	if len(interfaceListResp.Records) > 0 {
		for _, record := range interfaceListResp.Records {
			var tmp types.InterfaceInfo
			_ = copier.Copy(&tmp, record)
			list = append(list, tmp)
		}
	}
	return &types.PageListResp{
		Total:   interfaceListResp.Total,
		Records: list,
	}, nil
}
