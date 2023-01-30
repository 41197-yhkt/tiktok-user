package main

import (
	"gorm.io/gen"
	"tiktok-user/dao/dal/model"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,
	})

	g.ApplyBasic(model.User{}, model.UserRelation{})

	g.ApplyInterface(func(method model.UserMethod) {}, model.User{})
	g.ApplyInterface(func(method model.UserRelationMethod) {}, model.UserRelation{})

	g.Execute()
}
