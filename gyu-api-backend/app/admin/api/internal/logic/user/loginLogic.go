package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-backend/app/admin/api/internal/models"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/tools"
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
	if req.Username == "" || req.Password == "" || len(req.Username) < 6 || len(req.Password) < 8 {
		return nil, xerr.NewErrCodeMsg(xerr.RequestParamError, "用户名或密码错误")
	}
	_, err = regexp.MatchString(constant.PatternStr, req.Username)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(xerr.RequestParamError, "用户名称包含非法字符")
	}

	// 根据用户名和密码查询是否存在用户
	userModel := models.NewDefaultUserModel(l.svcCtx.DBEngin)
	user, err := userModel.SearchUserByUsername(req.Username)
	// 如果用户不存在，登陆失败，返回
	if err != nil {
		return nil, err
	}
	// 如果用户存在，再校验用户密码是否正确
	err = checkUserPassword(user.Password, req.Password)
	if err != nil {
		return nil, err
	}
	// 用户名和密码都无误且用户存在，生成 jwt token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&GenerateTokenReq{userId: user.Id})
	if err != nil {
		return nil, err
	}
	// token 存入缓存
	// key field value 的格式如下
	// login:token:xxx {userId: xxx, userRole: xxx, username: xxx, avatarUrl: xxx}
	tokenLogic := models.NewDefaultTokenModel(l.svcCtx.RedisClient)
	err = tokenLogic.InsertToken(tokenResp.accessToken, user.Id, user.UserRole, user.Username, user.AvatarUrl)
	if err != nil {
		return nil, err
	}

	// 登陆成功，返回用户 id，用户名，token，token 过期时间
	return &types.LoginResp{
		Id:          user.Id,
		Username:    user.Username,
		AvatarUrl:   user.AvatarUrl,
		UserRole:    user.UserRole,
		Token:       tokenResp.accessToken,
		TokenExpire: tokenResp.accessExpire,
	}, nil
}

// 校验用户密码

func checkUserPassword(pwd string, password string) error {
	str, err := tools.DecodeMd5(pwd)
	if err != nil {
		return xerr.NewErrCode(xerr.DecodeMd5Error)
	}
	if !tools.DecodeBcrypt(str, password) {
		return xerr.NewErrCode(xerr.UserPasswordError)
	}
	return nil
}
