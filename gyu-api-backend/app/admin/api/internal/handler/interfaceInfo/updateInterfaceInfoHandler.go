package interfaceInfo

import (
	"gyu-api-backend/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gyu-api-backend/app/admin/api/internal/logic/interfaceInfo"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
)

func UpdateInterfaceInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateInterfaceInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := interfaceInfo.NewUpdateInterfaceInfoLogic(r.Context(), svcCtx)
		resp, err := l.UpdateInterfaceInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
