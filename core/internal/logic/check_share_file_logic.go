package logic

import (
	"cloud-disk/core/db"
	"cloud-disk/core/define"
	"cloud-disk/core/process"
	"cloud-disk/core/utils"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"

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

func (l *CheckShareFileLogic) CheckShareFile(req *types.CheckShareFileRequest) (resp *types.CommonResponse, e error) {
	// key: ShareFileId
	// Value: FieldId, UserId, ExpiredTime, ClickNum
	fileId, userId, expiredAt, clickNum := uint(0), uint(0), int64(0), uint(0)
	key := define.ShareFilePrefix + fmt.Sprintf("%d", req.ShareFileId)
	resp = new(types.CommonResponse)

	// 从缓存中拿数据
	if res, err := l.svcCtx.RedisClient.Get(l.ctx, key).Result(); err != redis.Nil {
		a := strings.Split(res, ":")
		fileId, userId, expiredAt = utils.ToUint(a[0]), utils.ToUint(a[1]), utils.ToInt64(a[2])

		if expiredAt != -1 && expiredAt < time.Now().Unix() {
			resp.Code = 401
			resp.Msg = "分享文件已过期"
			return
		}

		clickNum = utils.ToUint(l.svcCtx.RedisClient.Get(l.ctx, key+"_ClickNum").Val())
		process.AddTask(func() {
			// 更新刷新数
			l.svcCtx.RedisClient.Set(l.ctx, key+"_ClickNum", fmt.Sprintf("%d", clickNum), 0)
			// 把点击次数保存到数据库
			if clickNum++; clickNum%define.UpdateClickNum == 0 {
				err := db.UpdateShareFieldClickNumById(req.ShareFileId, clickNum)
				if err != nil {
					utils.Logger().Error(err)
				}
			}
		})

		resp.Code = 200
		resp.Msg = "操作成功"
		resp.Data = utils.H{"file_id": fileId, "user_id": userId, "expired_at": expiredAt, "click_num": clickNum}
		return
	}

	// 从数据库中拿数据
	shareFile, err := db.GetShareFileById(req.ShareFileId)
	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		return
	}

	// 保存缓存
	process.AddTask(func() {
		val := fmt.Sprintf("%d:%d:%d", fileId, userId, expiredAt)
		l.svcCtx.RedisClient.Set(l.ctx, key, val, define.CacheExpireTime*time.Second)
		l.svcCtx.RedisClient.Set(l.ctx, key+"_ClickNum", fmt.Sprintf("%d", shareFile.ClickNum+1), 0)
	})

	// 返回结果
	fileId, userId, expiredAt, clickNum = shareFile.FileId, shareFile.UserId, shareFile.ExpiredTime, shareFile.ClickNum+1
	resp.Code = 200
	resp.Msg = "操作成功"
	resp.Data = utils.H{"file_id": fileId, "user_id": userId, "expired_at": expiredAt, "click_num": clickNum}
	return
}
