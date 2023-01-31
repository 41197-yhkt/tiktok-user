package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"tiktok-user/dao/dal"
	"tiktok-user/dao/dal/model"
	"tiktok-user/dao/dal/query"
	"tiktok-user/kitex_gen/user"
	"tiktok-user/util"
)

var q = query.Use(dal.DB.Debug())

func RegisterUser(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RegisterUser")
	defer span.Finish()

	rawPassword := req.Password
	password, err := util.EncryptPasswd(rawPassword)
	if err != nil {
		span.LogFields(
			log.String("EncryptPasswdError", err.Error()),
		)

		return nil, err
	}

	newUser := &model.User{
		Name:     req.Username,
		Password: password,
	}

	ud := q.User.WithContext(ctx)
	err = ud.Create(newUser)
}
