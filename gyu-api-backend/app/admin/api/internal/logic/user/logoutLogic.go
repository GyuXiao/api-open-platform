package user

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
	"gyu-api-backend/app/admin/rpc/client/user"
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
	logoutResp, err := l.svcCtx.UserRpc.Logout(l.ctx, &user.LogoutReq{AuthToken: token})
	if err != nil {
		return nil, err
	}

	return &types.LogoutResp{IsLogouted: logoutResp.IsLogouted}, nil
}
