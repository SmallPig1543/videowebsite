package dao

import (
	"context"
	"gorm.io/gorm"
	"videoweb/repository/db/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(user).Error
	return
}

func (dao *UserDao) FindUserByUserName(userName string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).
		First(&user).Error
	return
}

func (dao *UserDao) FindUserByUserEmail(email string) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).First(&user, "email=?", email).Error
	if err != nil {
		return nil, err
	}
	return
}

func (dao *UserDao) UpdateTotpStatus(user *model.User, ok bool) (err error) {
	err = dao.DB.Model(&user).Update("totp_enable_status", ok).Error
	return
}

func (dao *UserDao) UpdateOtpSecret(user *model.User, secret string, url string) (err error) {
	err = dao.DB.Model(&user).Update("totp_secret", secret).Error
	if err != nil {
		return err
	}
	err = dao.DB.Model(&user).Update("totp_url", url).Error
	return err
}

func (dao *UserDao) UpdateAvatar(user *model.User, avatarURL string) (err error) {
	err = dao.DB.Model(&user).Update("avatar_file_name", avatarURL).Error
	return
}
