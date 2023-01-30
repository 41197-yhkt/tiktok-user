package main

import (
	"context"
	user "tiktok-user/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoResponse) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// UserFollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserFollow(ctx context.Context, req *user.UserFollowRequest) (resp *user.UserFollowResponse, err error) {
	// TODO: Your code here...
	return
}

// UserUnfollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserUnfollow(ctx context.Context, req *user.UserFollowRequest) (resp *user.UserFollowResponse, err error) {
	// TODO: Your code here...
	return
}
