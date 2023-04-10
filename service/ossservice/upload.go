package ossservice

import (
	"bytes"
	"fmt"
	"mediafiledownload/config"
	"mediafiledownload/service/initservice"
)

func Upload(c config.OssConfig, mediabytes []byte) {
	bucket, err := initservice.Oss.Bucket("<yourBucketName>")
	if err != nil {
		fmt.Println("Error creating OSS bucket:", err)
		return
	}

	// Upload the file to the bucket with a unique key
	err = bucket.PutObject(c.UniqueKey, bytes.NewReader(mediabytes))
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}
}
