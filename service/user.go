package service

import (
	"bytes"
	"context"
	"errors"
	"gorm.io/gorm"
	"image/png"
	"os"
	"sync"
	"time"
	"videoweb/pkg/ctl"
	"videoweb/pkg/e"
	"videoweb/pkg/util"
	"videoweb/repository/db/dao"
	"videoweb/repository/db/model"
	"videoweb/types"
)

var UserServiceOnce sync.Once
var UserServiceIns *UserService

type UserService struct {
}

func GetUserService() *UserService {
	UserServiceOnce.Do(func() {
		UserServiceIns = &UserService{}
	})
	return UserServiceIns
}

func (s *UserService) Register(ctx context.Context, req *types.UserRegisterReq) (interface{}, error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(req.UserName)
	//没找到就创建新用户
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = &model.User{
			UserName: req.UserName,
			Email:    req.Email,
			CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		}
		//密码加密
		if err = user.SetPassword(req.Password); err != nil {
			util.LogrusObj.Info(err)
			code := e.SetPasswordFail
			return ctl.RespError(err, code), err
		}

		//存入数据库
		if err = userDao.CreateUser(user); err != nil {
			util.LogrusObj.Info(err)
			code := e.ErrorDataBase
			return ctl.RespError(err, code), err
		}
		return ctl.RespSuccess(), nil
	}
	//找到报错返回
	code := e.ErrorUserExist
	err = errors.New("user exists")
	return ctl.RespError(err, code), nil
}

func (s *UserService) EmailLogin(ctx context.Context, req *types.UserEmailLoginReq) (interface{}, error) {
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserEmail(req.Email)
	//没找到
	if err != nil {
		code := e.ErrorUserNotExist
		return ctl.RespError(err, code), err
	}
	//校验密码
	if !user.CheckPassword(req.Password) {
		err = errors.New("密码错误")
		util.LogrusObj.Info(err)
		code := e.ErrorPassword
		return ctl.RespError(err, code), err
	}

	if user.TotpEnableStatus {
		ok := util.VerifyOtp(req.OTP, user.TotpSecret)
		if !ok {
			code := e.VerifyOtpFailed
			err = errors.New("verify otp failed")
			return ctl.RespError(err, code), err
		}
	}
	//生成token
	token, err := util.GenerateToken(user.ID, user.UserName)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.TokenGeneratedFail
		return ctl.RespError(err, code), err
	}
	//返回数据
	u := &types.UserResp{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreateAt,
	}
	data := &types.TokenData{
		User:  u,
		Token: token,
	}
	return ctl.RespSuccessWithData(data), nil
}

func (s *UserService) EnableTotp(ctx context.Context, req *types.UserEnableTotpReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(u.UserName)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorDataBase
		return ctl.RespError(err, code), err
	}
	if user.TotpEnableStatus && req.Status == 1 {
		code := e.ERROR
		err = errors.New("you have enabled")
		return ctl.RespError(err, code), err
	}

	status := true
	if req.Status == 0 {
		ok := util.VerifyOtp(req.OTP, user.TotpSecret)
		if !ok {
			code := e.VerifyOtpFailed
			err = errors.New("verify otp failed")
			return ctl.RespError(err, code), err
		}
		status = false
	}

	err = userDao.UpdateTotpStatus(user, status)
	if err != nil {
		code := e.UpdateTotpStatusFailed
		return ctl.RespError(err, code), err
	}

	if status {
		key, err := util.GenerateOtp(user.UserName)
		if err != nil {
			util.LogrusObj.Info(err)
			code := e.ErrorGenerateOTP
			return ctl.RespError(err, code), err
		}
		err = userDao.UpdateOtpSecret(user, key.Secret(), key.URL())
		if err != nil {
			code := e.ErrorDataBase
			return ctl.RespError(err, code), err
		}
		img, _ := key.Image(200, 200)
		var buf bytes.Buffer
		_ = png.Encode(&buf, img)
		storePath := "./static/img/" + user.UserName + ".png"
		err = os.WriteFile(storePath, buf.Bytes(), 0644) // 把二维码写到文件里
		if err != nil {
			code := e.ERROR
			return ctl.RespError(err, code), err
		}
		err = util.SendEmail(user.Email, storePath) // 传入二维码的存储路径
		if err != nil {
			code := e.ErrorSendEmail
			return ctl.RespError(err, code), err
		}
		return ctl.RespSuccessWithData("Wait for your email"), err
	}
	return ctl.RespSuccess(), nil
}

func (s *UserService) AvatarUpload(ctx context.Context, req *types.UserAvatarUploadReq) (interface{}, error) {
	u, err := ctl.GetUserInfo(ctx)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorGetUserInfo
		return ctl.RespError(err, code), err
	}
	key, err := util.AvatarUpload(req.AvatarFileName)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorAvatarUpload
		return ctl.RespError(err, code), err
	}

	//更新db
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.FindUserByUserName(u.UserName)
	if err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorDataBase
		return ctl.RespError(err, code), err
	}
	if err = userDao.UpdateAvatar(user, key); err != nil {
		util.LogrusObj.Info(err)
		code := e.ErrorDataBase
		return ctl.RespError(err, code), err
	}
	return ctl.RespSuccess(), nil
}
