package initservice

import (
	"mediafiledownload/config"
	"mediafiledownload/service/logservice"
	"mediafiledownload/service/ossservice"
	"mediafiledownload/service/redisservice"
)

func InitService(c config.Config) {
	logservice.InitLogger(c.Log)
	Rdb = redisservice.InitRedis(c.Redis)
	Oss = ossservice.InitOss(c.Oss)
}
