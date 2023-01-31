// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "tiktok-user/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UserRegister(ctx context.Context, req *user.UserRegisterRequest, callOptions ...callopt.Option) (r *user.UserRegisterResponse, err error)
	UserLogin(ctx context.Context, req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error)
	UserInfo(ctx context.Context, req *user.UserInfoResponse, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error)
	UserFollow(ctx context.Context, req *user.UserFollowRequest, callOptions ...callopt.Option) (r *user.UserFollowResponse, err error)
	UserUnfollow(ctx context.Context, req *user.UserFollowRequest, callOptions ...callopt.Option) (r *user.UserFollowResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) UserRegister(ctx context.Context, req *user.UserRegisterRequest, callOptions ...callopt.Option) (r *user.UserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserRegister(ctx, req)
}

func (p *kUserServiceClient) UserLogin(ctx context.Context, req *user.UserLoginRequest, callOptions ...callopt.Option) (r *user.UserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserLogin(ctx, req)
}

func (p *kUserServiceClient) UserInfo(ctx context.Context, req *user.UserInfoResponse, callOptions ...callopt.Option) (r *user.UserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserInfo(ctx, req)
}

func (p *kUserServiceClient) UserFollow(ctx context.Context, req *user.UserFollowRequest, callOptions ...callopt.Option) (r *user.UserFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserFollow(ctx, req)
}

func (p *kUserServiceClient) UserUnfollow(ctx context.Context, req *user.UserFollowRequest, callOptions ...callopt.Option) (r *user.UserFollowResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserUnfollow(ctx, req)
}