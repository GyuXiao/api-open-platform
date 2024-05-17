package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-interface/user/internal/svc"
	"gyu-api-interface/user/internal/types"
)

type GetUsernameByPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsernameByPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsernameByPostLogic {
	return &GetUsernameByPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsernameByPostLogic) GetUsernameByPost(req *types.PostUserReq) (resp *types.PostUserResp, err error) {
	// 成功响应
	// 测试一下是否会造成请求超时
	// time.Sleep(time.Minute * 10)
	logc.Infof(l.ctx, "GetUsernameByPost requestBody: %s", req.Username)
	return &types.PostUserResp{PostResp: "GetUsernameByPost response: " + req.Username}, nil
}
