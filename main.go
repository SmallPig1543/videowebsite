package main

import (
	"videoweb/config"
	"videoweb/pkg/util"
	"videoweb/repository/cache"
	"videoweb/repository/db/dao"
	"videoweb/route"
)

func main() {
	r := route.NewRouter()
	_ = r.Run(":9090")
}

func init() {
	config.InitConfig()
	util.OssInit()
	util.InitLog()
	dao.InitMySQL()
	cache.LinkRedis()
}
