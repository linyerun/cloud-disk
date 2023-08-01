package handler

import (
	"cloud-disk/core/db"
	"cloud-disk/core/utils"
	"io/ioutil"
	"net/http"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadFileByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析参数
		var req types.DownloadFileByIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取文件信息
		resp := new(types.CommonResponse)
		file, err := db.GetFileById(req.FileId)
		if err != nil {
			resp.Code = 500
			resp.Msg = err.Error()
			httpx.OkJsonCtx(r.Context(), w, resp)
			return
		}

		// 下载文件
		cosFile, err := utils.TencentDownloadFile(file.Path)
		if err != nil {
			resp.Code = 500
			resp.Msg = err.Error()
			httpx.OkJsonCtx(r.Context(), w, resp)
			return
		}
		defer cosFile.Body.Close()

		// 返回文件
		bytes, err := ioutil.ReadAll(cosFile.Body)
		if err != nil {
			resp.Code = 500
			resp.Msg = err.Error()
			httpx.OkJsonCtx(r.Context(), w, resp)
			return
		}
		w.Header().Set("Content-Type", cosFile.Response.Header.Get("Content-Type"))
		w.Write(bytes)

	}
}
