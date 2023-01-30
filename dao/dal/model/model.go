package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model           // ID uint CreatAt time.Time UpdateAt time.Time DeleteAt gorm.DeleteAt If it is repeated with the definition will be ignored
	ID            uint   `gorm:"primary_key"`
	Name          string `gorm:"column: user_name"`
	Password      string `gorm:"column: user_pwd_hash"`
	FollowCount   int    `gorm:"column: follow_count; type: int"`
	FollowerCount int    `gorm:"column: follower_count; type: int"`
}

type UserRelation struct {
	gorm.Model      // ID uint CreatAt time.Time UpdateAt time.Time DeleteAt gorm.DeleteAt If it is repeated with the definition will be ignored
	ID         uint `gorm:"primary_key"`
	FollowFrom uint `gorm:"column: follow_from; type:bigint"`
	FollowTo   uint `gorm:"column: follow_to; type:bigint"`
}
