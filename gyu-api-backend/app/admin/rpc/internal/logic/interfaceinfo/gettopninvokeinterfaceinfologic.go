package interfaceinfologic

import (
	"context"
	"github.com/jinzhu/copier"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/common/constant"

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
	userInterfaceInfoModel := models.NewDefaultUserInterfaceInfoModel(l.svcCtx.DBEngin)
	if in.Limit == 0 {
		in.Limit = constant.DefaultTopNLimit
	}
	results, err := userInterfaceInfoModel.GetTopInvokeInterfaceInfoList(in.Limit)
	if err != nil {
		return nil, err
	}

	var records []*pb.InvokeInterfaceInfo
	for _, itf := range results {
		tmp := &pb.InvokeInterfaceInfo{}
		_ = copier.Copy(tmp, itf)
		records = append(records, tmp)
	}

	return &pb.GetTopNInvokeInterfaceInfoResp{Records: records}, nil
}
