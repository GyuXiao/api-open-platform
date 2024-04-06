package user

import (
	"context"
	"gyu-api-backend/app/admin/api/internal/models"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/tools"
	"gyu-api-backend/common/xerr"
	"regexp"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 校验参数
	if req.Username == "" || req.Password == "" || len(req.Username) < 6 || len(req.Password) < 8 || req.Password != req.ConfirmPassword {
		return nil, xerr.NewErrCodeMsg(xerr.RequestParamError, "用户名或密码错误")
	}
	_, err = regexp.MatchString(constant.PatternStr, req.Username)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(xerr.RequestParamError, "用户名称包含非法字符")
	}

	// 先通过 username 查询用户是否存在
	userModel := models.NewDefaultUserModel(l.svcCtx.DBEngin)
	user, err := userModel.SearchUserByUsername(req.Username)
	// 如果存在，返回用户已经存在，注册失败
	if user != nil {
		return nil, xerr.NewErrCode(xerr.UserExistError)
	}
	if err != nil && err.(*xerr.CodeError).GetErrCode() != xerr.RecordNotFoundError {
		return nil, err
	}
	// 用户第一次注册，调用 createUser 创建用户
	// 处于数据安全考虑，用户密码存入数据库前先做加密处理
	pwd, pwdErr := encodeUserPassword(req.Password)
	if pwdErr != nil {
		return nil, pwdErr
	}
	userMap := map[string]interface{}{
		"username": req.Username,
		"password": pwd,
	}
	err = userModel.CreateUser(userMap)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.CreateUserError)
	}

	// 注册成功，返回用户名
	return &types.RegisterResp{Username: req.Username}, nil
}

// 用户密码加密

func encodeUserPassword(pwd string) (string, error) {
	hashStr, err := tools.EncodeBcrypt(pwd)
	if err != nil {
		return "", xerr.NewErrCode(xerr.EncryptionError)
	}
	return tools.EncodeMd5([]byte(hashStr)), nil
}
