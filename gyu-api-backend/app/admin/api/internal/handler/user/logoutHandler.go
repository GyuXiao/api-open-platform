package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"gyu-api-backend/app/admin/api/internal/types"
	"gyu-api-backend/common/result"
	"net/http"

	"gyu-api-backend/app/admin/api/internal/logic/user"
	"gyu-api-backend/app/admin/api/internal/svc"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogoutReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := user.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout(&req)
		result.HttpResult(r, w, resp, err)
	}
}
