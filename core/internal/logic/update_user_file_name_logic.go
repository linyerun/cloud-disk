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

type UpdateUserFileNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserFileNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserFileNameLogic {
	return &UpdateUserFileNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserFileNameLogic) UpdateUserFileName(req *types.UpdateUserFileNameRequest) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	defer func() {
		if len(resp.Msg) == 0 {
			resp.Msg = resp_code_msg.GetMsgByCode(resp.Code)
		}
	}()
	// 校验参数
	if !utils.IsAllowLen(req.Filename, 1, 255) {
		resp.Code = resp_code_msg.ParamsError
		resp.Msg = "文件名长度太长"
		return
	}
	// 执行更新操作
	err = db.UpdateUserFileName(req.UserId, req.UserFileId, req.Filename)
	if err != nil {
		resp.Code = 501
		resp.Msg = err.Error()
		return
	}
	// 成功
	resp.Code = resp_code_msg.Success
	return
}
