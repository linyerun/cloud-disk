// Code generated by goctl. DO NOT EDIT.
package types

type CommonResponse struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type SendCodeRequest struct {
	Email string `json:"email"`
}

type RegisterRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Nickname     string `json:"nickname"`
	HeadPortrait string `json:"head_portrait"`
	Code         string `json:"code"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}