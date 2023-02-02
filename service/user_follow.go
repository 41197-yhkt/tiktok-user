package service

import (
	"context"
	"github.com/41197-yhkt/tiktok-user/dao/dal"
	"github.com/41197-yhkt/tiktok-user/dao/dal/model"
	"github.com/41197-yhkt/tiktok-user/dao/dal/query"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/util"
	"github.com/opentracing/opentracing-go"
)

// UserFollow 用户关注，就算关注多次，数据库实际也只保存一条记录昂
func UserFollow(ctx context.Context, req *user.UserFollowRequest) (resp *user.UserFollowResponse, err error) {
	resp = user.NewUserFollowResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserInfo")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)

	followFrom := uint(req.GetFollowFrom())
	followTo := uint(req.GetFollowTo())

	newUserRelation := model.UserRelation{
		FollowFrom: followFrom,
		FollowTo:   followTo,
	}

	err = userRelationDao.Create(&newUserRelation)
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}

	resp.BaseResp = util.PackBaseResp(nil)
	return

}

// UserUnfollow 用户取消关注，由于存在软删除
func UserUnfollow(ctx context.Context, req *user.UserUnfollowRequest) (resp *user.UserUnfollowResponse, err error) {
	resp = user.NewUserUnfollowResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserInfo")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)

	followFrom := uint(req.GetFollowFrom())
	followTo := uint(req.GetFollowTo())

	// 存在这条记录才需要删除
	newUserRelation := model.UserRelation{
		FollowFrom: followFrom,
		FollowTo:   followTo,
	}

	_, err = userRelationDao.Delete(&newUserRelation)
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}

	resp.BaseResp = util.PackBaseResp(nil)
	return
}
