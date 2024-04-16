package interfaceInfo

import (
	"gyu-api-backend/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gyu-api-backend/app/admin/api/internal/logic/interfaceInfo"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
)

func GetInterfaceInfoByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInterfaceInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := interfaceInfo.NewGetInterfaceInfoByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetInterfaceInfoById(&req)
		result.HttpResult(r, w, resp, err)
	}
}
