package model

import "gorm.io/gen"

type UserMethod interface {
	//where("user_name=@user_name")
	FindByUserName(user_name string) (gen.T, error)
	//where("id=@id")
	FindByUserID(id int64) (gen.T, error)
	//update @@table
	//	{{set}}
	//		update_time=now(),
	//		{{if follow_count > 0}}
	//			follow_count=@follow_count
	//		{{end}}
	//	{{end}}
	// where id=@id
	UpdateUserFollowCount(id int, follow_count int) error
	//update @@table
	//	{{set}}
	//		update_time=now(),
	//		{{if follower_count != ""}}
	//			follower_count=@follower_count
	//		{{end}}
	//	{{end}}
	// where id=@id
	UpdateUserFollowerCount(id int, follower_count string) error
}

type UserRelationMethod interface {
	//where("id=@id")
	FindByID(id int64) (gen.T, error)
	//where("follow_from=@follow_from")
	FindByFollowFrom(follow_from string) ([]gen.T, error)
	//where("follow_to=@follow_to")
	FindByFollowTo(follow_to string) ([]gen.T, error)
	//where("follow_from=@follow_from and follow_to=@follow_to")
	FindByFollowFromAndFollowTo(follow_from string, follow_to string) (gen.T, error)
}
