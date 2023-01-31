// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "tiktok-user/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

var userServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
	methods := map[string]kitex.MethodInfo{
		"UserRegister": kitex.NewMethodInfo(userRegisterHandler, newUserServiceUserRegisterArgs, newUserServiceUserRegisterResult, false),
		"UserLogin":    kitex.NewMethodInfo(userLoginHandler, newUserServiceUserLoginArgs, newUserServiceUserLoginResult, false),
		"UserInfo":     kitex.NewMethodInfo(userInfoHandler, newUserServiceUserInfoArgs, newUserServiceUserInfoResult, false),
		"UserFollow":   kitex.NewMethodInfo(userFollowHandler, newUserServiceUserFollowArgs, newUserServiceUserFollowResult, false),
		"UserUnfollow": kitex.NewMethodInfo(userUnfollowHandler, newUserServiceUserUnfollowArgs, newUserServiceUserUnfollowResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserRegisterArgs)
	realResult := result.(*user.UserServiceUserRegisterResult)
	success, err := handler.(user.UserService).UserRegister(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserRegisterArgs() interface{} {
	return user.NewUserServiceUserRegisterArgs()
}

func newUserServiceUserRegisterResult() interface{} {
	return user.NewUserServiceUserRegisterResult()
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserLoginArgs)
	realResult := result.(*user.UserServiceUserLoginResult)
	success, err := handler.(user.UserService).UserLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserLoginArgs() interface{} {
	return user.NewUserServiceUserLoginArgs()
}

func newUserServiceUserLoginResult() interface{} {
	return user.NewUserServiceUserLoginResult()
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserInfoArgs)
	realResult := result.(*user.UserServiceUserInfoResult)
	success, err := handler.(user.UserService).UserInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserInfoArgs() interface{} {
	return user.NewUserServiceUserInfoArgs()
}

func newUserServiceUserInfoResult() interface{} {
	return user.NewUserServiceUserInfoResult()
}

func userFollowHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserFollowArgs)
	realResult := result.(*user.UserServiceUserFollowResult)
	success, err := handler.(user.UserService).UserFollow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserFollowArgs() interface{} {
	return user.NewUserServiceUserFollowArgs()
}

func newUserServiceUserFollowResult() interface{} {
	return user.NewUserServiceUserFollowResult()
}

func userUnfollowHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUserUnfollowArgs)
	realResult := result.(*user.UserServiceUserUnfollowResult)
	success, err := handler.(user.UserService).UserUnfollow(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUserUnfollowArgs() interface{} {
	return user.NewUserServiceUserUnfollowArgs()
}

func newUserServiceUserUnfollowResult() interface{} {
	return user.NewUserServiceUserUnfollowResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (r *user.UserRegisterResponse, err error) {
	var _args user.UserServiceUserRegisterArgs
	_args.Req = req
	var _result user.UserServiceUserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, req *user.UserLoginRequest) (r *user.UserLoginResponse, err error) {
	var _args user.UserServiceUserLoginArgs
	_args.Req = req
	var _result user.UserServiceUserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, req *user.UserInfoResponse) (r *user.UserInfoResponse, err error) {
	var _args user.UserServiceUserInfoArgs
	_args.Req = req
	var _result user.UserServiceUserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserFollow(ctx context.Context, req *user.UserFollowRequest) (r *user.UserFollowResponse, err error) {
	var _args user.UserServiceUserFollowArgs
	_args.Req = req
	var _result user.UserServiceUserFollowResult
	if err = p.c.Call(ctx, "UserFollow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserUnfollow(ctx context.Context, req *user.UserFollowRequest) (r *user.UserFollowResponse, err error) {
	var _args user.UserServiceUserUnfollowArgs
	_args.Req = req
	var _result user.UserServiceUserUnfollowResult
	if err = p.c.Call(ctx, "UserUnfollow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}