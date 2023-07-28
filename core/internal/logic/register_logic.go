package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/define"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.CommonResponse, e error) {
	resp = new(types.CommonResponse)
	defer func() { resp.Msg = resp_code_msg.GetMsgByCode(resp.Code) }()
	// 校验参数
	if !utils.IsAllowPwd(req.Password) || !utils.IsNormalEmail(req.Email) || !utils.IsAllowLen(req.Nickname, 1, 40) {
		resp.Code = resp_code_msg.ParamsError
		return
	}
	// 如果nickname重复, 给nickname一个随机值
	if db.HasTheNickname(req.Nickname) {
		req.Nickname += utils.RandomCode(10)
	}
	// 判断验证码是否合法
	if err := l.svcCtx.RedisClient.Get(l.ctx, "register_"+req.Email+"_"+req.Code).Err(); err != nil && err.Error() == db.RedisNil {
		resp.Code = resp_code_msg.ParamsError
		return
	}
	// 保存用户信息
	user := &models.User{
		Email:        req.Email,
		Password:     utils.Md5(req.Password),
		Nickname:     req.Nickname,
		HeadPortrait: req.HeadPortrait,
	}
	if err := db.SaveUser(user); err != nil {
		resp.Code = resp_code_msg.SaveDataError
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
