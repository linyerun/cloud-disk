package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/dto"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFileListLogic {
	return &GetUserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserFileListLogic) GetUserFileList(req *types.GetUserFileListRequest) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	defer func() {
		if len(resp.Msg) == 0 {
			resp.Msg = resp_code_msg.GetMsgByCode(resp.Code)
		}
	}()
	userFiles, err := db.GetUserFileListByParentIdUserId(req.ParentId, req.UserId)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		utils.Logger().Error(err)
		return
	}
	userFilesDto := make([]*dto.UserFile, len(userFiles))
	for i := range userFiles {
		userFilesDto[i] = &dto.UserFile{
			ID:       userFiles[i].ID,
			FileId:   userFiles[i].FileId,
			Filename: userFiles[i].Filename,
			FileType: userFiles[i].FileType,
			ParentId: req.ParentId,
		}
	}
	resp.Code = resp_code_msg.Success
	resp.Data = utils.H{"user_files": userFilesDto}
	return
}
