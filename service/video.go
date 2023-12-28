package service

import (
	"context"
	"sync"
	"time"
	"videoweb/pkg/ctl"
	"videoweb/pkg/e"
	"videoweb/pkg/util"
	"videoweb/repository/db/dao"
	"videoweb/repository/db/model"
	"videoweb/types"
)

var VideoServiceOnce sync.Once
var VideoServiceIns *VideoService

type VideoService struct {
}

func GetVideoService() *VideoService {
	VideoServiceOnce.Do(func() {
		VideoServiceIns = &VideoService{}
	})
	return VideoServiceIns
}
func (s *VideoService) VideoCreate(ctx context.Context, req *types.VideoCreateRequest) (interface{}, error) {
	key, err := util.VideoUpload(req.LocalFileName)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorVideoUpload
		return ctl.RespError(err, code), err
	}
	v := &model.Video{
		Uid:      0,
		Title:    req.Title,
		Types:    req.Types,
		Key:      key,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	videoDao := dao.NewVideoDao(ctx)
	err = videoDao.CreateUser(v)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorDataBase
		return ctl.RespError(err, code), err
	}
	url, err := util.GetURL(v.Key)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUrl
		return ctl.RespError(err, code), err
	}
	resp := &types.VideoResp{
		ID:       v.ID,
		Uid:      v.Uid,
		Title:    v.Title,
		Types:    v.Types,
		URL:      url,
		Views:    0,
		CreateAt: v.CreateAt,
	}
	return ctl.RespSuccessWithData(resp), nil
}
