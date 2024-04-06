package user

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/models"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"strconv"
	"strings"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentLogic {
	return &CurrentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentLogic) Current(req *types.CurrentUserReq) (resp *types.CurrentUserResp, err error) {
	token := strings.Split(req.Authorization, " ")[1]
	// 校验从 jwt token 解析出来的 userId 是否和缓存中的 userId 一致
	// 1，从 token 中解析出 userId1
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	claims, err := generateTokenLogic.ParseTokenByKey(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return nil, err
	}
	userId1 := claims[constant.KeyJwtUserId].(float64)
	// 2，从 根据 token 从 redis 中拿到 userId2
	tokenLogic := models.NewDefaultTokenModel(l.svcCtx.RedisClient)
	result, err := tokenLogic.CheckTokenExist(token)
	if err != nil {
		return nil, err
	}
	userIdStr, userRoleStr, username, avatarUrl := result[0], result[1], result[2], result[3]
	userId2, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, err
	}
	userRole, err := strconv.Atoi(userRoleStr)
	if err != nil {
		return nil, err
	}
	// 3，判断两者是否相同
	if uint64(userId1) != uint64(userId2) {
		return nil, xerr.NewErrCode(xerr.UserNotLoginError)
	}
	// 校验成功后，刷新 token
	tokenLogic.RefreshToken(token)
	return &types.CurrentUserResp{
		Id:          uint64(userId1),
		Username:    username,
		AvatarUrl:   avatarUrl,
		UserRole:    uint8(userRole),
		Token:       token,
		TokenExpire: int64(constant.TokenExpireTime.Seconds()),
	}, nil
}
