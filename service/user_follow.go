package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"tiktok-user/dao/dal"
	"tiktok-user/dao/dal/model"
	"tiktok-user/dao/dal/query"
	"tiktok-user/kitex_gen/user"
	"tiktok-user/util"
)

// UserFollow 用户关注，就算关注多次，数据库实际也只保存一条记录昂
func UserFollow(ctx context.Context, req *user.UserFollowRequest) (resp *user.UserFollowResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserInfo")
	defer span.Finish()

	var q = query.Use(dal.DB.Debug())
	userRelationDao := q.UserRelation.WithContext(ctx)

	followFrom := uint(req.GetFollowFrom())
	followTo := uint(req.GetFollowTo())
	userRelation := &model.UserRelation{}

	db := dal.DB
	queryRes := db.Unscoped().Where("follow_from = ?", followFrom).Where("follow_to = ?", followTo).Take(&userRelation)
	err = queryRes.Error

	// 存在这条记录
	if err == nil && queryRes.RowsAffected == 1 {
		err = db.Unscoped().Model(&userRelation).Update("deleted_at", nil).Error
		if err != nil {
			resp.BaseResp = util.PackBaseResp(err)
			return
		}

		resp.BaseResp = util.PackBaseResp(nil)
		return
	} else {
		// 不存在这条记录，就插入一条呗
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
}

// UserUnfollow 用户取消关注，由于存在软删除
func UserUnfollow(ctx context.Context, req *user.UserUnfollowRequest) (resp *user.UserUnfollowResponse, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserInfo")
	defer span.Finish()

	followFrom := uint(req.GetFollowFrom())
	followTo := uint(req.GetFollowTo())
	userRelation := &model.UserRelation{}

	db := dal.DB
	queryRes := db.Unscoped().Where("follow_from = ?", followFrom).Where("follow_to = ?", followTo).Take(&userRelation)
	err = queryRes.Error

	// 存在这条记录才需要删除
	if err == nil && queryRes.RowsAffected == 1 {
		err = db.Unscoped().Delete(&queryRes).Error
		if err != nil {
			resp.BaseResp = util.PackBaseResp(err)
			return
		}
	}

	resp.BaseResp = util.PackBaseResp(nil)
	return
}
