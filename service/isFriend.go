package service

import (
	"context"

	"github.com/41197-yhkt/tiktok-user/dao/dal"
	"github.com/41197-yhkt/tiktok-user/dao/dal/query"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/util"
	"github.com/opentracing/opentracing-go"
)

func IsFriend(ctx context.Context, userId, toUserId int64) (resp *user.IsFriendResponse, err error) {
	resp = user.NewIsFriendResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "IsFriend")
	defer span.Finish()
	var q = query.Use(dal.DB.Debug())
	UserRelationDao := q.UserRelation.WithContext(ctx)

	//找到user1的关注列表
	userId_follow_list, err1 := UserRelationDao.FindByFollowFrom(uint(userId))
	if err1 != nil {
		resp.BaseResp = util.PackBaseResp(err1)
		return
	}
	//找到user2的关注列表
	toUserId_follow_list, err2 := UserRelationDao.FindByFollowFrom(uint(toUserId))
	if err2 != nil {
		resp.BaseResp = util.PackBaseResp(err2)
		return
	}
	var flag1, flag2 bool
	//判断user1的关注列表是否有user2
	for _, v := range userId_follow_list {
		if toUserId == int64(v.FollowTo) {
			flag1 = true
			break
		}
	}
	//判断user2的关注列表是否有user1
	for _, v := range toUserId_follow_list {
		if userId == int64(v.FollowTo) {
			flag2 = true
			break
		}
	}
	resp.IsFriend = flag1 && flag2
	resp.BaseResp = util.PackBaseResp(nil)
	return resp, nil
}
