package handler

import (
	"net/http"

	"service/api/rest/internal/svc"
	"service/api/rest/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func (h *Handler) CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// l := logic.NewCreateUserLogic(r.Context(), svcCtx)
		err := h.cul.CreateUser(r.Context(), &req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
