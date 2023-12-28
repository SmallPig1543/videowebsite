package e

var MsgFlags = map[int]string{
	SUCCESS: "操作成功",
	ERROR:   "操作失败",

	InvalidParams: "请求参数有误",

	ErrorUserExist:    "用户已存在",
	ErrorUserNotExist: "用户不存在",

	//token
	TokenGeneratedFail: "token生成失败",
	ErrorTokenTimeout:  "token超时",

	//user
	SetPasswordFail:        "密码设置失败",
	ErrorGetUserInfo:       "获取用户信息失败",
	VerifyOtpFailed:        "otp验证失败",
	UpdateTotpStatusFailed: "更新2FA状态失败",
	ErrorGenerateOTP:       "生成otp失败",
	ErrorPassword:          "密码错误",
	ErrorGetAvatar:         "获取头像失败",
	ErrorAvatarUpload:      "头像上传失败",

	//video
	ErrorVideoOpen:   "视频打开失败",
	ErrorVideoUpload: "视频上传失败",
	ErrorGetUrl:      "获取视频url失败",

	//email
	ErrorSendEmail: "发送邮件失败",
	ErrorEmailAuth: "邮件鉴权有误",
}

// GetMsg 获取错误码对应的信息
func GetMsg(code int) string {
	if msg, ok := MsgFlags[code]; ok {
		return msg
	}
	return MsgFlags[ERROR]
}
