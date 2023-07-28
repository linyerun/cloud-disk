package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/define"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"context"
	"time"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeForRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeForRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeForRegisterLogic {
	return &SendCodeForRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeForRegisterLogic) SendCodeForRegister(req *types.SendCodeRequest) (resp *types.CommonResponse, e error) {
	resp = new(types.CommonResponse)
	defer func() { resp.Msg = resp_code_msg.GetMsgByCode(resp.Code) }()
	// 判断邮箱是否合法
	// 判断这个邮箱是否被注册过了
	if !utils.IsNormalEmail(req.Email) || db.HasRegistered(req.Email) {
		resp.Code = resp_code_msg.ParamsError
		return
	}
	// 判断邮箱短期又没发送邮件
	// 存在, 那么Get的时候不会报错
	if err := l.svcCtx.RedisClient.Get(l.ctx, define.RedisCodePrefix+req.Email).Err(); err == nil {
		resp.Code = resp_code_msg.EmailSendManyError
		utils.Logger().Println("用户:" + req.Email + ", 打算频繁发送验证码")
		return
	}
	// 生成验证码并发送到对应邮箱
	code := utils.GetCode()
	err := utils.SendCode(req.Email, code)
	if err != nil {
		resp.Code = resp_code_msg.SendEmailError
		return
	}
	// 设置验证码过期时间
	l.svcCtx.RedisClient.Set(l.ctx, "register_"+req.Email+"_"+code, "", define.EachCodeForEmailWaitTime*time.Second)
	// 避免邮箱短时间内再次发送
	l.svcCtx.RedisClient.Set(l.ctx, define.RedisCodePrefix+req.Email, "", define.EachCodeForEmailWaitTime*time.Second)
	// 发送成功
	resp.Code = resp_code_msg.Success
	return
}
