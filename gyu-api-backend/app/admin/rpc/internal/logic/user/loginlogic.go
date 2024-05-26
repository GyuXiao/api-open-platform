package userlogic

import (
	"context"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/common/tools"
	"gyu-api-backend/common/xerr"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// 根据用户名和密码查询是否存在用户
	userModel := models.NewDefaultUserModel(l.svcCtx.DBEngin)
	user, err := userModel.SearchUserByUsername(in.Username)
	// 如果用户不存在，登陆失败，返回
	if err != nil {
		return nil, err
	}
	// 如果用户存在，再校验用户密码是否正确
	err = checkUserPassword(user.Password, in.Password)
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
	return &pb.LoginResp{
		Id:          user.Id,
		Username:    user.Username,
		AvatarUrl:   user.AvatarUrl,
		UserRole:    uint64(user.UserRole),
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
