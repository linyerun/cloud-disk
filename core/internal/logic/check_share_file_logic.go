package logic

import (
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckShareFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckShareFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckShareFileLogic {
	return &CheckShareFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckShareFileLogic) CheckShareFile(req *types.CheckShareFileRequest) (resp *types.CommonResponse, err error) {
	// todo: add your logic here and delete this line
	return
}
