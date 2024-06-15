package result

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"gyu-api-backend/common/xerr"
	"net/http"
)

type JsonResponse struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		successResp := Success(resp)
		httpx.WriteJson(w, http.StatusOK, successResp)
		return
	}

	code := xerr.ServerCommonError
	msg := xerr.GetMsgByCode(code)

	causeErr := errors.Cause(err)
	e, ok := causeErr.(*xerr.CodeError)
	if ok {
		code = e.GetErrCode()
		msg = e.GetErrMsg()
	} else {
		grpcErr, rpcOk := status.FromError(causeErr)
		if rpcOk {
			grpcCode := uint32(grpcErr.Code())
			if xerr.IsCodeErr(grpcCode) {
				code = grpcCode
				msg = grpcErr.Message()
			}
		}
	}

	logc.Errorf(r.Context(), "【API-ERR】 : %+v ", err)
	httpx.WriteJson(w, http.StatusBadRequest, Error(code, msg))
}

func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	logc.Errorf(context.Background(), "HTTP 请求参数错误: %v", err.Error())
	errMsg := fmt.Sprintf("%s： %v", xerr.GetMsgByCode(xerr.RequestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(xerr.ParamFormatError, errMsg))
}
