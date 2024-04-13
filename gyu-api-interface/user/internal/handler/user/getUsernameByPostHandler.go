package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gyu-api-interface/user/internal/logic/user"
	"gyu-api-interface/user/internal/svc"
	"gyu-api-interface/user/internal/types"
)

func GetUsernameByPostHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := user.NewGetUsernameByPostLogic(r.Context(), svcCtx)
		resp, err := l.GetUsernameByPost(&req, r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
