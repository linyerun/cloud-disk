package handler

import (
	"net/http"
	"strconv"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveShareFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveShareFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 加入token解析处理的userId
		userIdStr := r.Header.Get("user_id")
		uid, _ := strconv.Atoi(userIdStr)
		req.UserId = uint(uid)

		l := logic.NewSaveShareFileLogic(r.Context(), svcCtx)
		resp, err := l.SaveShareFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
