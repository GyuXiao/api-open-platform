package userlogic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/common/xerr"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInvokeUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetInvokeUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInvokeUserLogic {
	return &GetInvokeUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetInvokeUserLogic) GetInvokeUser(in *pb.GetInvokeUserReq) (*pb.GetInvokeUserResp, error) {
	if in.AccessKey == "" {
		logc.Info(l.ctx, "accessKey 不能为空")
		return nil, xerr.NewErrCode(xerr.AccessKeyNotExistError)
	}
	// 根据 accessKey 检索用户信息
	userModel := models.NewDefaultUserModel(l.svcCtx.DBEngin)
	user, err := userModel.SearchUserByAccessKey(in.AccessKey)
	if err != nil {
		return nil, err
	}

	return &pb.GetInvokeUserResp{
		Id:        user.Id,
		Username:  user.Username,
		Password:  user.Password,
		AvatarUrl: user.AvatarUrl,
		Email:     user.Email,
		Phone:     user.Phone,
		UserRole:  uint64(user.UserRole),
		AccessKey: user.AccessKey,
		SecretKey: user.SecretKey,
	}, nil
}
