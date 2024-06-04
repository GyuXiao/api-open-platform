package analysis

import (
	"context"
	"github.com/jinzhu/copier"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/common/userTools"
	"strings"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopNInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTopNInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopNInterfaceInfoLogic {
	return &GetTopNInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTopNInterfaceInfoLogic) GetTopNInterfaceInfo(req *types.GetTopNInterfaceInfoReq) (resp *types.GetTopNInterfaceInfoResp, err error) {
	// 1,校验是否为管理员
	token := strings.Split(req.Authorization, " ")[1]
	err = userTools.CheckUserIsAdminRole(l.svcCtx.RedisClient, token)
	if err != nil {
		return nil, err
	}

	// 2，获取 topN 的调用接口信息
	topResp, err := l.svcCtx.InterfaceInfoRpc.GetTopNInvokeInterfaceInfo(l.ctx, &interfaceinfo.GetTopNInvokeInterfaceInfoReq{
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}

	var records []types.InvokeInterfaceInfo
	if len(topResp.Records) > 0 {
		for _, record := range topResp.Records {
			var tmp types.InvokeInterfaceInfo
			_ = copier.Copy(&tmp, record)
			records = append(records, tmp)
		}
	}
	return &types.GetTopNInterfaceInfoResp{Records: records}, nil
}
