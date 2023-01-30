gen_service:
	kitex -module "tiktok-user" -service tiktok-user ../idl/user.thrift

gen_gorm:
	cd ./dao/cmd/generate/ && go run .