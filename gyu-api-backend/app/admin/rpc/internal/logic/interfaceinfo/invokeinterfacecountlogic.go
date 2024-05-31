package interfaceinfologic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"gyu-api-backend/app/admin/models"
	"gyu-api-backend/common/xerr"

	"gyu-api-backend/app/admin/rpc/internal/svc"
	"gyu-api-backend/app/admin/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InvokeInterfaceCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvokeInterfaceCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvokeInterfaceCountLogic {
	return &InvokeInterfaceCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvokeInterfaceCountLogic) InvokeInterfaceCount(in *pb.UpdateInvokeInterfaceCountReq) (*pb.UpdateInvokeInterfaceCountResp, error) {
	if in.UserId == 0 || in.InterfaceInfoId == 0 {
		logc.Info(l.ctx, "更新接口调用次数时 userId 或者 interfaceInfoId 不能为空")
		return nil, xerr.NewErrCode(xerr.RequestParamError)
	}

	userInterfaceInfoModel := models.NewDefaultUserInterfaceInfoModel(l.svcCtx.DBEngin)
	err := userInterfaceInfoModel.UpdateForInvokeSuccess(in.UserId, in.InterfaceInfoId)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateInvokeInterfaceCountResp{IsUpdated: true}, nil
}
