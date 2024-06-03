package interfaceInfo

import (
	"context"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/app/admin/rpc/client/interfaceinfo"
	"gyu-api-backend/common/constant"
	"gyu-api-backend/common/xerr"
	"strconv"
	"strings"

	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnlineInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnlineInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineInterfaceInfoLogic {
	return &OnlineInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnlineInterfaceInfoLogic) OnlineInterfaceInfo(req *types.OnlineInterfaceInfoReq) (resp *types.OnlineInterfaceInfoResp, err error) {
	// 0 通过 token 获取 redis 存储的 userRole，如果不是管理者，则不能执行上线操作
	token := strings.Split(req.Authorization, " ")[1]
	tokenLogic := models.NewDefaultTokenModel(l.svcCtx.RedisClient)
	result, err := tokenLogic.CheckTokenExist(token)
	if err != nil {
		return nil, err
	}
	userRoleStr := result[1]
	username := result[2]
	userRole, _ := strconv.Atoi(userRoleStr)
	if userRole != constant.AdminRole {
		return nil, xerr.NewErrCode(xerr.PermissionDenied)
	}

	onlineInterfaceInfoResp, err := l.svcCtx.InterfaceInfoRpc.OnlineInterfaceInfo(l.ctx, &interfaceinfo.OnlineInterfaceInfoReq{
		Id:       req.Id,
		Username: username,
	})
	if err != nil {
		return nil, err
	}

	return &types.OnlineInterfaceInfoResp{IsOnline: onlineInterfaceInfoResp.IsOnline}, nil
}
