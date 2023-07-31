package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserFileByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserFileByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserFileByIdLogic {
	return &DeleteUserFileByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserFileByIdLogic) DeleteUserFileById(req *types.DeleteUserFileByIdRequest) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	err = db.DeleteUserFileById(req.UserFileId, req.UserId)
	if err != nil {
		resp.Code = resp_code_msg.SystemError
		resp.Msg = err.Error()
		utils.Logger().Error(err)
		return
	}
	resp.Code = resp_code_msg.Success
	resp.Msg = resp_code_msg.GetMsgByCode(resp.Code)
	return
}
