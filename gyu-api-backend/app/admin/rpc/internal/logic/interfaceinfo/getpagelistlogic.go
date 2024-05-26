package interfaceinfologic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.PageListResp{}, nil
}
