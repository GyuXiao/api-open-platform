package userlogic

import (
	"context"
	"gyu-api-backend/app/admin/models"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *pb.LogoutReq) (*pb.LogoutResp, error) {
	tokenLogic := models.NewDefaultTokenModel(l.svcCtx.RedisClient)
	err := tokenLogic.DeleteToken(in.AuthToken)
	if err != nil {
		return nil, err
	}
	return &pb.LogoutResp{IsLogouted: true}, nil
}
