package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"videoweb/api"
	"videoweb/middleware"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api")
	{
		user := v1.Group("user")
		{
			user.POST("register", api.UserRegisterHandler())
			user.POST("avatar-upload", api.UserAvatarUploadHandler())
			user.POST("email-login", api.UserEmailLoginHandler())
			user.POST("enable-totp", middleware.JWT(), api.UserEnableTotpHandler())
		}
		video := v1.Group("video")
		{
			video.POST("create", api.VideoCreateHandler())

		}
	}
	return r
}
