package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"videoweb/pkg/util"
	"videoweb/service"
	"videoweb/types"
)

func VideoCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.VideoCreateRequest
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetVideoService()
			resp, err := l.VideoCreate(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
			return
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
