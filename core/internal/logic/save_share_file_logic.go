package logic

import (
	"context"

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
	// todo: add your logic here and delete this line
	return
}
