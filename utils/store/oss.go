package store

import (
	"errors"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego"
)

var (
	endpoint        string = ""
	accessKeyId     string = ""
	accessKeySecret string = ""
	bucket          string = ""
)

func InitOss() {
	endpoint = beego.AppConfig.String("oss_endpoint")
	accessKeyId = beego.AppConfig.String("oss_access_key_id")
	accessKeySecret = beego.AppConfig.String("oss_access_key_secret")
	bucket = beego.AppConfig.String("oss_bucket")
}

func getOssBucket() (*oss.Bucket, error) {
	if len(endpoint) == 0 || len(accessKeyId) == 0 || len(accessKeySecret) == 0 {
		beego.Error("oss param error")
		return nil, errors.New("oss param error")
	}
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if nil != err {
		beego.Error("oss init error")
		return nil, errors.New("oss init error")
	}

	if len(bucket) == 0 {
		return nil, errors.New("get bucket error")
	}
	return client.Bucket(bucket)
}

func OssPutObject(ossPath, localFilePath string) error {
	bucket, err := getOssBucket()
	if nil != err {
		return err
	}

	err = bucket.PutObjectFromFile(ossPath, localFilePath)
	if nil != err {
		beego.Error(err.Error())
	}
	return err
}
