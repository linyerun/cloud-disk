package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadFileByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadFileByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadFileByIdLogic {
	return &DownloadFileByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadFileByIdLogic) DownloadFileById(req *types.DownloadFileByIdRequest) (resp *types.CommonResponse, err error) {
	return
}
