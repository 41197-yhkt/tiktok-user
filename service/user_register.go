package service

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"tiktok-user/dao/dal"
	"tiktok-user/dao/dal/model"
	"tiktok-user/dao/dal/query"
	"tiktok-user/kitex_gen/user"
	"tiktok-user/util"
)

func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRegister")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userDao := q.User.WithContext(ctx)
	_, err = userDao.FindByUserName(req.Username)

	if err == nil {
		// username对应的用户已经存在，请重新注册
		sErr := errors.New("username is already in use, please change another one")
		resp.BaseResp = util.PackBaseResp(sErr)
		return resp, sErr
	}

	// 当前注册的用户已经不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 对于用户密码进行加密处理
		rawPassword := req.Password
		password, err := util.EncryptPasswd(rawPassword)
		if err != nil {
			resp.BaseResp = util.PackBaseResp(err)
			return resp, err
		}

		newUser := &model.User{
			Name:     req.Username,
			Password: password,
		}

		// 创建新用户失败
		createRes := dal.DB.Create(newUser)
		if createRes.Error != nil {
			resp.BaseResp = util.PackBaseResp(createRes.Error)
			return resp, createRes.Error
		}

		// 正常返回
		resp.UserId = int64(newUser.ID)
		resp.BaseResp = util.PackBaseResp(nil)
		return resp, nil
	} else {
		resp.BaseResp = util.PackBaseResp(err)
		return resp, err
	}
}
