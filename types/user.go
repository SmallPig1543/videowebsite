package types

type UserResp struct {
	ID        uint   `json:"id" form:"id"`
	UserName  string `json:"user_name" form:"user_name" `
	CreateAt  string `json:"create_at" form:"create_at"`
	Email     string `json:"email" form:"email"`
	AvatarURL string `json:"avatar_url" form:"avatar_url"`
}

type UserRegisterReq struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password,required" form:"password"`
	Email    string `json:"email" form:"email,required"`
}

type UserEmailLoginReq struct {
	Email    string `json:"email" form:"email,required"`
	Password string `json:"password" form:"password,required"`
	OTP      string `json:"otp" form:"otp"`
}

type UserAvatarUploadReq struct {
	AvatarFileName string `json:"avatar_file_name" form:"avatar_file_name"` //本地文件的完整路径
}

type UserEnableTotpReq struct {
	// Status 0关闭,1开启
	Status int    `json:"status" form:"status"`
	OTP    string `json:"otp" form:"otp"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
