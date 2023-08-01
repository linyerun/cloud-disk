package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/define"
	"cloud-disk/core/models"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"context"
	"fmt"
	"time"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveShareFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveShareFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveShareFileLogic {
	return &SaveShareFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveShareFileLogic) SaveShareFile(req *types.SaveShareFileRequest) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	if req.ExpiredTime != -1 && req.ExpiredTime <= time.Now().Unix() {
		resp.Code = resp_code_msg.ParamsError
		resp.Msg = "过期时间设置存在问题"
		return
	}

	fileId, err := db.GetUserFileFileIdByIds(req.UserId, req.UserFileId)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		utils.Logger().Error(err)
		return
	}

	shareFile := &models.ShareFile{
		UserId:      req.UserId,
		FileId:      fileId,
		ExpiredTime: req.ExpiredTime,
	}
	if err = db.SaveShareFile(shareFile); err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		utils.Logger().Error(err)
		return
	}

	resp.Code = resp_code_msg.Success
	resp.Msg = "操作成功"
	resp.Data = utils.H{"share_url": define.Host + "/share/file/check/" + fmt.Sprintf("%d", shareFile.ID), "method": "GET"}
	return
}
