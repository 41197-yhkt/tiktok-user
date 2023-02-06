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

func GetFriendList(ctx context.Context, req *user.FriendListRequest) (resp *user.FriendListResponse, err error) {
	resp = user.NewFriendListResponse()
	span, ctx := opentracing.StartSpanFromContext(ctx, "FriendList")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)
	userDao := q.User.WithContext(ctx)
	userId := req.UserId

	//找到user的关注列表，得到的userFollow_list的follow_from均为userId
	userFollow_list, err := userRelationDao.FindByFollowFrom(uint(userId))

	if err != nil {
		resp.BaseResp = util.PackBaseResp(err)
		return
	}

	for _, v := range userFollow_list {
		//判断user的关注列表有没有跟user是朋友的
		bool_resp, erro := IsFriend(ctx, userId, int64(v.FollowTo))
		if erro != nil {
			resp.BaseResp = util.PackBaseResp(err)
			return
		}
		//如果是朋友，则直接找出该follow_user信息，并加入到resp user_list中
		if bool_resp.IsFriend {
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
	}
	resp.BaseResp = util.PackBaseResp(nil)
	return resp, nil
}
