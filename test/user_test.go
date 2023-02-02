package test

import (
	"context"
	"fmt"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/service"
	"testing"
)

func TestRegister(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserRegisterRequest()
	req.Username = "lyh_test"
	req.Password = "123456"
	resp, err := service.UserRegister(ctx, req)
	fmt.Println("TestRegister resp = ", resp, " err = ", err)
}

func TestReregister(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserRegisterRequest()
	req.Username = "lyh_test"
	req.Password = "123456"
	resp, err := service.UserRegister(ctx, req)
	fmt.Println("TestRegister resp = ", resp, " err = ", err)
}

func TestLoginSuccess(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserLoginRequest()
	req.Username = "lyh_test"
	req.Password = "123456"
	resp, err := service.UserLogin(ctx, req)
	fmt.Println("TestLoginSuccess resp = ", resp, " err = ", err)
}

func TestLoginFailedWrongPassword(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserLoginRequest()
	req.Username = "lyh_test"
	req.Password = "12345678"
	resp, err := service.UserLogin(ctx, req)
	fmt.Println("TestLoginFailed resp = ", resp, " err = ", err)
}

func TestLoginFailedWrongUsername(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserLoginRequest()
	req.Username = "lyh_test_2"
	req.Password = "123456"
	resp, err := service.UserLogin(ctx, req)
	fmt.Println("TestLoginFailed resp = ", resp, " err = ", err)
}

func TestUserInfo(t *testing.T) {
	ctx := context.Background()
	req := user.NewUserLoginRequest()
	req.Username = "lyh_test"
	req.Password = "123456"
	resp, err := service.UserLogin(ctx, req)
	fmt.Println("TestLoginSuccess resp = ", resp, " err = ", err)

	existenceUserId := resp.UserId
	infoReq := user.NewUserInfoRequest()
	infoReq.UserId = existenceUserId
	userInfo, err := service.UserInfo(ctx, infoReq)
	fmt.Println("TestUserInfo resp = ", userInfo, " err = ", err)
}
