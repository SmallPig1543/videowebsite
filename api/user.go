package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"videoweb/pkg/util"
	"videoweb/service"
	"videoweb/types"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetUserService()
			resp, err := l.Register(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func UserEnableTotpHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserEnableTotpReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetUserService()
			resp, err := l.EnableTotp(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
func UserEmailLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserEmailLoginReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetUserService()
			resp, err := l.EmailLogin(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}

func UserAvatarUploadHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserAvatarUploadReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetUserService()
			resp, err := l.AvatarUpload(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
