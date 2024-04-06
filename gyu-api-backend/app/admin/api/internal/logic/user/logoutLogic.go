package user

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/models"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.LogoutReq) (resp *types.LogoutResp, err error) {
	token := strings.Split(req.Authorization, " ")[1]
	tokenLogic := models.NewDefaultTokenModel(l.svcCtx.RedisClient)
	err = tokenLogic.DeleteToken(token)
	if err != nil {
		return nil, err
	}
	return &types.LogoutResp{IsLogouted: true}, nil
}
