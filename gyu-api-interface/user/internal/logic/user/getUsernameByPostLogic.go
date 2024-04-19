package user

import (
	"context"
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gyu-api-interface/user/internal/svc"
	"gyu-api-interface/user/internal/types"
	"net/http"
)

var AccessKey string = "gyu"
var SecretKey string = "abcdefg"

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

func (l *GetUsernameByPostLogic) GetUsernameByPost(req *types.PostUserReq, r *http.Request) (resp *types.PostUserResp, err error) {
	// API 校验逻辑
	headers := r.Header
	accessKey := headers.Get("AccessKey")
	nonce := headers.Get("Nonce")
	timestamp := headers.Get("Timestamp")
	sign := headers.Get("Sign")
	body := headers.Get("Body")
	// accessKey 校验
	if accessKey != AccessKey {
		logc.Info(l.ctx, "accessKey 错误, 权限校验未通过")
		return nil, xerror.New("accessKey 错误, 权限校验未通过")
	}
	paramsMap := map[string]string{
		"title0": AccessKey,
		"title1": SecretKey,
		"title2": nonce,
		"title3": timestamp,
		"title4": body,
	}
	signature := GenSign(paramsMap, SecretKey)
	// 签名校验
	if signature != sign {
		logc.Info(l.ctx, "签名错误, 权限校验未通过")
		return nil, xerror.New("签名错误, 权限校验未通过")
	}
	// 成功响应
	return &types.PostUserResp{PostResp: "GetUsernameByPost response: " + req.Username}, nil
}

func GenSign(params map[string]string, secretKey string) string {
	concatString := concatMapString(params)
	hms := cryptor.HmacSha256(concatString, secretKey)
	return hms
}

func concatMapString(paramsMap map[string]string) string {
	n := len(paramsMap)
	sortKeys := make([]string, n)
	for i := 0; i < n; i++ {
		sortKeys[i] = fmt.Sprintf("title%d", i)
	}
	concatString := ""
	for _, key := range sortKeys {
		concatString += paramsMap[key]
	}
	return concatString
}
