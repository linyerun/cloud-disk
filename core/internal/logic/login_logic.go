package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/define"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.CommonResponse, e error) {
	resp = new(types.CommonResponse)
	defer func() { resp.Msg = resp_code_msg.GetMsgByCode(resp.Code) }()
	// 校验参数
	if !utils.IsNormalEmail(req.Email) || !utils.IsAllowPwd(req.Password) {
		resp.Code = resp_code_msg.ParamsError
		return
	}
	// 判断是否能登录
	if !db.HasTheUser(req.Email, utils.Md5(req.Password)) {
		resp.Code = resp_code_msg.LoginError
		return
	}
	// 生成token
	token, err := utils.GenerateToken(req.Email, define.TokenExpireTime)
	refreshToken, err := utils.GenerateToken(req.Email, define.TokenExpireTime*2)
	if err != nil {
		resp.Code = resp_code_msg.TokenGenerateError
		return
	}
	// 返回token
	resp.Code = resp_code_msg.Success
	resp.Data = utils.H{"token": token, "refresh_token": refreshToken}
	return
}
