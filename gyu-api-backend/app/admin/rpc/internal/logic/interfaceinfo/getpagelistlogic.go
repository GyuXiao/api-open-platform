package interfaceinfologic

import (
	"context"
	"github.com/jinzhu/copier"
	"gyu-api-backend/app/admin/models"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPageListLogic {
	return &GetPageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPageListLogic) GetPageList(in *pb.PageListReq) (*pb.PageListResp, error) {
	interfaceInfoLogic := models.NewDefaultInterfaceInfoModel(l.svcCtx.DBEngin)
	result, total, err := interfaceInfoLogic.FindListPage(in.Keyword, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}
	var interfaceInfoList []*pb.InterfaceInfo
	for _, interfaceInfo := range result {
		tmp := &pb.InterfaceInfo{}
		_ = copier.Copy(tmp, interfaceInfo)
		interfaceInfoList = append(interfaceInfoList, tmp)
	}
	return &pb.PageListResp{
		Total:   uint64(total),
		Records: interfaceInfoList,
	}, nil
}
