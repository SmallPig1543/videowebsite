package dao

import (
	"context"
	"gorm.io/gorm"
	"videoweb/repository/db/model"
)

type VideoDao struct {
	*gorm.DB
}

func NewVideoDao(ctx context.Context) *VideoDao {
	if ctx == nil {
		ctx = context.Background()
	}

	return &VideoDao{NewDBClient(ctx)}
}

func (dao *VideoDao) CreateUser(video *model.Video) (err error) {
	err = dao.DB.Model(&model.Video{}).Create(video).Error
	return
}
