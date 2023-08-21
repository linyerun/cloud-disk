package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshTokenLogic {
	return &RefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshTokenLogic) RefreshToken(req *types.UserIdRequest) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	defer func() {
		if resp.Msg == "" {
			resp.Msg = resp_code_msg.GetMsgByCode(resp.Code)
		}
	}()
	// 生成token
	token, err := utils.GenerateToken(req.Email, req.UserId, define.TokenExpireTime)
	refreshToken, err := utils.GenerateToken(req.Email, req.UserId, define.TokenExpireTime*2)
	if err != nil {
		resp.Code = resp_code_msg.TokenGenerateError
		return
	}
	// 返回token
	resp.Code = resp_code_msg.Success
	resp.Data = utils.H{"token": token, "refresh_token": refreshToken}
	return
}
