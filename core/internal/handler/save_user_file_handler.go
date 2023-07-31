package handler

import (
	"net/http"
	"strconv"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveUserFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveUserFile
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 把userId加入请求参数中
		userIdStr := r.Header.Get("user_id")
		uid, _ := strconv.Atoi(userIdStr)
		req.UserId = uint(uid)
		// 调用逻辑层
		l := logic.NewSaveUserFileLogic(r.Context(), svcCtx)
		resp, err := l.SaveUserFile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
