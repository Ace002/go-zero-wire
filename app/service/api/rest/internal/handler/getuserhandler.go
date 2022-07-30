package handler

import (
	"net/http"

	"service/api/rest/internal/svc"
	"service/api/rest/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func (h *Handler) GetUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		// l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := h.gul.GetUser(r.Context(), &req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
