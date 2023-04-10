package ossservice

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"mediafiledownload/config"
)

func InitOss(c config.OssConfig) *oss.Client {
	client, err := oss.New(c.AccessKeyId, c.AccessKeySecret, c.Endpoint)
	if err != nil {
		zap.L().Error(fmt.Sprintf("init oss service failed,err:%v", err.Error()))
		panic("init oss service failed")
	}

	return client

}
