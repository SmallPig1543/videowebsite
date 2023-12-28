package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"videoweb/pkg/ctl"
	"videoweb/pkg/e"
	"videoweb/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(http.StatusForbidden, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "缺少Token",
			})
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ErrorTokenFail //无权限
		} else if time.Now().Unix() > claims.ExpiresAt.Unix() {
			code = e.ErrorTokenTimeout
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "token有误",
			})
			c.Abort()
			return
		}
		//创建新ctx.request
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{ID: claims.ID, UserName: claims.UserName}))
		c.Next()
	}
}
