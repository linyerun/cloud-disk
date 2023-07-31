package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/resp_code_msg"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserFileParentIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserFileParentIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserFileParentIdLogic {
	return &UpdateUserFileParentIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserFileParentIdLogic) UpdateUserFileParentId(req *types.UpdateUserFileParentIdRequest) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	if !db.HasTheDir(req.ParentId, req.UserId) {
		resp.Code = resp_code_msg.ParamsError
		resp.Msg = "文件夹不存在"
		return
	}
	err = db.UpdateUserFileParentId(req.UserId, req.UserFileId, req.ParentId)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		return
	}
	resp.Code = resp_code_msg.Success
	resp.Msg = resp_code_msg.GetMsgByCode(resp.Code)
	return
}
