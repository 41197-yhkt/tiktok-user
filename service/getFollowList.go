package service

import (
	"context"
	"errors"

	"github.com/41197-yhkt/pkg/errno"
	"github.com/41197-yhkt/tiktok-user/dao/dal"
	"github.com/41197-yhkt/tiktok-user/dao/dal/query"
	"github.com/41197-yhkt/tiktok-user/kitex_gen/user"
	"github.com/41197-yhkt/tiktok-user/util"
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

func GetFollowList(ctx context.Context, req *user.FollowListRequest) (resp *user.FollowListResponse, err error) {
	resp = user.NewFollowListResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "FollowList")
	defer span.Finish()
	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)
	userDao := q.User.WithContext(ctx)
	userID := req.UserId
	//找到user的关注列表
	userRelation_list, err := userRelationDao.FindByFollowFrom(uint(userID))
	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}
	//根据user的关注列表ID找到对应UserInfo
	for _, v := range userRelation_list {
		follow_user, sErr := userDao.FindByUserID(v.FollowTo)
		if sErr != nil {
			if errors.Is(sErr, gorm.ErrRecordNotFound) {
				resp.BaseResp = util.PackBaseResp(errno.UserNotExist)
			} else {
				resp.BaseResp = util.PackBaseResp(sErr)
			}
			return resp, sErr
		}
		follow_count := int64(follow_user.FollowCount)
		follower_count := int64(follow_user.FollowerCount)
		resp.UserList = append(resp.UserList, &user.User{
			Id:            int64(follow_user.ID),
			Name:          follow_user.Name,
			FollowCount:   &follow_count,
			FollowerCount: &follower_count,
			IsFollow:      false,
		})
	}
	resp.BaseResp = util.PackBaseResp(nil)
	return resp, nil

}
