package initservice

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
)

var (
	Rdb *redis.Client
	Oss *oss.Client
)
