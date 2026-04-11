// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package handler

import (
	"net/http"

	"github.com/cylixlee/go-framework-playground/internal/logic"
	"github.com/cylixlee/go-framework-playground/internal/svc"
	"github.com/cylixlee/go-framework-playground/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignUpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignUpRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSignUpLogic(r.Context(), svcCtx)
		resp, err := l.SignUp(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
