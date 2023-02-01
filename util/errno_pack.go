package util

import (
	"github.com/41197-yhkt/pkg/errno"
	"tiktok-user/kitex_gen/user"
)

func PackBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	// 业务错误，更新错误信息
	s := errno.ServerError
	s.Msg = err.Error()
	return baseResp(s)
}

func baseResp(err *errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		StatusCode: int32(err.Code),
		StatusMsg:  &err.Msg,
	}
}
