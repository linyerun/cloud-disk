package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/models"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFile) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	defer func() { resp.Msg = resp_code_msg.GetMsgByCode(resp.Code) }()
	file := &models.File{
		Hash: req.Hash,
		Size: req.Size,
		Path: req.Path,
	}
	err = db.SaveFile(file)
	if err != nil {
		resp.Code = resp_code_msg.SaveDataError
		return
	}
	resp.Data = utils.H{"file_id": file.ID}
	return
}
