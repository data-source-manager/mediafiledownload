package initservice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"mediafiledownload/config"
)

type (
	Xredis interface {
		Subscribe(key string)
	}

	redisOp struct {
		rdb  *redis.Client
		conf config.RedisConfig
	}

	media struct {
		MediaType   string `json:"media_type"`
		MediaSrc    string `json:"media_src"`
		MediaMd5Src string `json:"media_md5_src"`
	}
)

func NewReid() Xredis {
	return &redisOp{rdb: Rdb}
}

func (r *redisOp) Subscribe(key string) {
	fmt.Println(key)
	sub := r.rdb.Subscribe(context.Background(), key)
	_, err := sub.Receive(context.Background())
	if err != nil {
		zap.L().Error(fmt.Sprintf("redis 订阅 error:%s", err.Error()))
		panic(err)
	}

	ch := sub.Channel()
	for msg := range ch {
		var m media
		err := json.Unmarshal([]byte(msg.Payload), &m)
		if err != nil {
			zap.L().Error(fmt.Sprintf("反序列化msg error:%s", msg.Payload))
			continue
		}

	}

}
