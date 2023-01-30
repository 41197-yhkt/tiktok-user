package dal

import (
	"fmt"
	"sync"
	"tiktok-user/dao/dal/model"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB
var once sync.Once

func init() {
	once.Do(func() {
		DB = ConnectDB().Debug()
		_ = DB.AutoMigrate(&model.User{}, &model.UserRelation{})
	})
}

func ConnectDB() (conn *gorm.DB) {
	dsn := "root:liu@tcp(127.0.0.1:3306)/tiktok-test?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	return conn
}
