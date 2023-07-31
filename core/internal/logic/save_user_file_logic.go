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

type SaveUserFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveUserFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserFileLogic {
	return &SaveUserFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveUserFileLogic) SaveUserFile(req *types.SaveUserFile) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	defer func() {
		if len(resp.Msg) == 0 {
			resp.Msg = resp_code_msg.GetMsgByCode(resp.Code)
		}
	}()
	// 校验参数
	if req.FileId != 0 && req.FileType == 0 {
		resp.Code = resp_code_msg.DirError
		return
	} else if req.FileId == 0 && req.FileType == 1 {
		resp.Code = resp_code_msg.FileError
		return
	} else if !utils.IsAllowLen(req.Filename, 0, 255) {
		resp.Code = resp_code_msg.ParamLenError
		return
	}
	// 校验parent_id是否合理
	if !db.HasTheDir(req.ParentId, req.UserId) {
		resp.Code = resp_code_msg.ParamsError
		resp.Msg = "不存在该父级目录"
		return
	}
	// 保存起来
	userFile := &models.UserFile{
		UserId:   req.UserId,
		ParentId: req.ParentId,
		FileId:   req.FileId,
		FileType: req.FileType,
		Filename: req.Filename,
	}
	err = db.SaveUserFile(userFile)
	if err != nil {
		resp.Code = resp_code_msg.SaveDataError
		resp.Msg = err.Error()
		return
	}
	// 返回结果
	resp.Code = resp_code_msg.Success
	resp.Data = utils.H{"user_file_id": userFile.ID}
	return
}
