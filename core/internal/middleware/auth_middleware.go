package middleware

import (
	"cloud-disk/core/internal/types"
	"cloud-disk/core/resp_code_msg"
	"cloud-disk/core/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		userClaim, err := utils.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			encoder := json.NewEncoder(w)
			if err := encoder.Encode(&types.CommonResponse{Code: resp_code_msg.TokenError, Msg: resp_code_msg.GetMsgByCode(resp_code_msg.TokenError)}); err != nil {
				utils.Logger().Errorln("{AuthMiddleware Handle}encoder.Encode err:", err)
				http.Error(w, err.Error(), 500)
			}
		}
		// 保存用户信息在header上面
		r.Header.Set("user_id", fmt.Sprintf("%d", userClaim.ID))
		r.Header.Set("user_email", userClaim.Email)
		// 放行
		next(w, r)
	}
}
