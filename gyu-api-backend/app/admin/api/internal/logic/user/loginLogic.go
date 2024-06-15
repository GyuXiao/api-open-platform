package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
	"gyu-api-backend/app/admin/rpc/client/user"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"regexp"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 校验参数
	if req.Username == constant.BlankString || req.Password == constant.BlankString || len(req.Username) < constant.UsernameMinLen || len(req.Password) < constant.PasswordMinLen {
		return nil, xerr.NewErrCodeMsg(xerr.RequestParamError, "用户名或密码错误")
	}
	_, err = regexp.MatchString(constant.PatternStr, req.Username)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(xerr.RequestParamError, "用户名称包含非法字符")
	}

	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Id:          loginResp.Id,
		Username:    loginResp.Username,
		AvatarUrl:   loginResp.AvatarUrl,
		UserRole:    uint8(loginResp.UserRole),
		Token:       loginResp.Token,
		TokenExpire: loginResp.TokenExpire,
	}, nil
}
