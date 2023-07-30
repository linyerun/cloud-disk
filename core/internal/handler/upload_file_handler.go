package handler

import (
	"bytes"
	"cloud-disk/core/db"
	"cloud-disk/core/define"
	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/utils"
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UploadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取文件和文件信息
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		} else if r.ContentLength > define.MaxFileSize {
			httpx.ErrorCtx(r.Context(), w, errors.New("文件超出500MB"))
			return
		}

		// 把文件转成bytes
		bs, err := ioutil.ReadAll(file)
		file.Close() // 关闭文件
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 判断用户是否可以上传这个文件
		userIdStr := r.Header.Get("user_id")
		uid, _ := strconv.Atoi(userIdStr)
		user, err := db.GetUserCapacityById(uid)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		} else if user.TotalCapacity < user.CurCapacity+uint(len(bs)) {
			httpx.ErrorCtx(r.Context(), w, errors.New("您的空间不足"))
			return
		}

		// 根据bs获取它的md5值
		hash := utils.Hash(bs)

		// 判断这个文件是否存在于File中了
		fid, err := db.GetFileByHash(hash)
		if err == nil {
			httpx.OkJsonCtx(r.Context(), w, types.CommonResponse{Code: 200, Msg: "上传成功", Data: utils.H{"file_id": fid}})
			return
		}

		// 进行文件上传操作
		fileURL, err := utils.TencentUploadFile(bytes.NewReader(bs), fileHeader.Filename)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 封装请求参数
		var req types.UploadFile
		req.Path = fileURL
		req.Hash = hash
		req.Size = uint(len(bs))

		// 调用逻辑层对应处理函数
		l := logic.NewUploadFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadFile(&req)

		// 返回结果到客户端
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
