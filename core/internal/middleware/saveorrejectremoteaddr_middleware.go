package middleware

import (
	"cloud-disk/core/internal/types"
	"cloud-disk/core/utils"
	"encoding/json"
	"net/http"
)

// IpBlackListSet 自定义黑名单
var IpBlackListSet = map[string]struct{}{}

type SaveOrRejectRemoteAddrMiddleware struct {
}

func NewSaveOrRejectRemoteAddrMiddleware() *SaveOrRejectRemoteAddrMiddleware {
	return &SaveOrRejectRemoteAddrMiddleware{}
}

func (m *SaveOrRejectRemoteAddrMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := r.Header.Get("x_forwarded_for")
		if len(remoteAddr) == 0 {
			remoteAddr = r.RemoteAddr
		}
		// 记录远程过来的IP地址
		utils.IpLogger().Infoln(remoteAddr)
		// 如果IP在黑名单, 拒绝请求
		if _, ok := IpBlackListSet[remoteAddr]; ok {
			w.WriteHeader(401)
			w.Header().Set("Content-Type", "application/json")
			encoder := json.NewEncoder(w)
			if err := encoder.Encode(&types.CommonResponse{Code: 401, Msg: "你已被加入黑名单，请联系管理员进行处理。"}); err != nil {
				utils.Logger().Errorln("{SaveOrRejectRemoteAddr_middleware Handle}encoder.Encode err:", err)
				http.Error(w, err.Error(), 500)
			}
			return
		}
		next(w, r)
	}
}
