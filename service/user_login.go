package service

import (
	"context"
	"errors"
	"github.com/41197-yhkt/pkg/errno"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"tiktok-user/dao/dal"
	"tiktok-user/dao/dal/query"
	"tiktok-user/kitex_gen/user"
	"tiktok-user/util"
)

func UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserLogin")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userDao := q.User.WithContext(ctx)
	gormUser, err := userDao.FindByUserName(req.Username)

	// 如果记录不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp.BaseResp = util.PackBaseResp(errno.UserNotExist)
		return resp, err
	}

	// 如果记录存在
	pwdCmpPass, err := util.ComparePasswd(req.Password, gormUser.Password)
	// 密码比对出现问题
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return resp, err
	}
	// 密码不匹配
	if !pwdCmpPass {
		resp.BaseResp = util.PackBaseResp(errno.UserPwdErr)
		return resp, err
	}

	resp.BaseResp = util.PackBaseResp(nil)
	return resp, nil
}
