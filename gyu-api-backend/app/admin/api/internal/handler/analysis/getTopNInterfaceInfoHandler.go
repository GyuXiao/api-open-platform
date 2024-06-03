package analysis

import (
	"gyu-api-backend/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gyu-api-backend/app/admin/api/internal/logic/analysis"
	"gyu-api-backend/app/admin/api/internal/svc"
	"gyu-api-backend/app/admin/api/internal/types"
)

func GetTopNInterfaceInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTopNInterfaceInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := analysis.NewGetTopNInterfaceInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetTopNInterfaceInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}
