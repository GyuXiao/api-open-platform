package interfaceinfologic

import (
	"context"
	"github.com/jinzhu/copier"
	"gyu-api-backend/app/admin/models"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopNInvokeInterfaceInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopNInvokeInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopNInvokeInterfaceInfoLogic {
	return &GetTopNInvokeInterfaceInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTopNInvokeInterfaceInfoLogic) GetTopNInvokeInterfaceInfo(in *pb.GetTopNInvokeInterfaceInfoReq) (*pb.GetTopNInvokeInterfaceInfoResp, error) {
	// 1，先查询到 topN 的 []{interfaceInfoId, totalNum}
	userInterfaceInfoModel := models.NewDefaultUserInterfaceInfoModel(l.svcCtx.DBEngin)
	results, err := userInterfaceInfoModel.GetTopInvokeInterfaceInfoList(in.Limit)
	if err != nil {
		return nil, err
	}
	interfaceInfoLogic := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)

	// 2，根据 interfaceInfoId 查询到 interfaceInfoName
	var list []*models.InvokeInterfaceInfoModel
	for _, res := range results {
		interfaceInfoId := res.InterfaceInfoId
		total := res.TotalNum
		interfaceInfo, err := interfaceInfoLogic.SearchInterfaceInfoById(interfaceInfoId)
		if err != nil {
			return nil, err
		}

		list = append(list, &models.InvokeInterfaceInfoModel{
			InterfaceInfoName: interfaceInfo.Name,
			TotalNum:          total,
		})
	}

	var records []*pb.InvokeInterfaceInfo
	for _, itf := range list {
		tmp := &pb.InvokeInterfaceInfo{}
		_ = copier.Copy(tmp, itf)
		records = append(records, tmp)
	}

	return &pb.GetTopNInvokeInterfaceInfoResp{Records: records}, nil
}
