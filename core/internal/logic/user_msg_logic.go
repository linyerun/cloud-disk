package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/dto"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMsgLogic {
	return &UserMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMsgLogic) UserMsg(req *types.UserMsgRequest) (resp *types.CommonResponse, err error) {
	resp = new(types.CommonResponse)
	defer func() { resp.Msg = resp_code_msg.GetMsgByCode(resp.Code) }()
	// 校验参数
	email := req.Email
	if !utils.IsNormalEmail(email) {
		resp.Code = resp_code_msg.ParamsError
		return
	}
	// 从数据库查找
	user, err := db.GetUserByEmail(email)
	if err != nil {
		resp.Code = resp_code_msg.GetDataError
		return
	}
	// 返回数据
	userMsgDto := &dto.User{
		Email:        user.Email,
		Nickname:     user.Nickname,
		HeadPortrait: user.HeadPortrait,
	}
	resp.Code = resp_code_msg.Success
	resp.Data = utils.H{"user_msg": userMsgDto}
	return
}
