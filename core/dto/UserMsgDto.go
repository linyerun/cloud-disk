package dto

type User struct {
	Email        string `json:"email"`
	Nickname     string `json:"nickname"`
	HeadPortrait string `json:"head_portrait,omitempty"`
}
