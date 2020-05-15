package service

import (
	"gin-user-center/app/config"
	"io"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OssClient struct {
	Client *oss.Client
}

var Oss *OssClient

func Init() {
	config := config.Conf
	client, err := oss.New(config.Oss.Endpoint, config.Oss.AccessKeyId, config.Oss.AccessKeySecret)
	if err != nil {
		logger.Error("oss init error:", err)
		os.Exit(-1)
	}
	Oss := &OssClient{}
	Oss.Client = client
}

/**
 * 创建bucket
 */
func (o *OssClient) CreateBucket(name string) error {
	return o.Client.CreateBucket(name)
}

/**
 * 删除bucket
 */
func (o *OssClient) DeleteBucket(name string) error {
	return o.Client.DeleteBucket(name)
}

/**
 * 上传
 */
func (o *OssClient) Upload(bucketName, key string, reader io.Reader) (err error) {
	// 获取存储空间。
	bucket, err := o.Client.Bucket(bucketName)
	if err != nil {
		logger.Error("util Upload Bucket error:", err)
		return err
	}

	// 简单上传
	if err = bucket.PutObject(key, reader); err != nil {
		logger.Error("util Upload PutObject error:", err)
		return err
	}

	return nil
}

func (o *OssClient) GetSignUrl(bucketName, key string) (signURL string, err error) {
	bucket, err := o.Client.Bucket(bucketName)
	if err != nil {
		logger.Error("util GetSignUrl Bucket error:", err)
		return "", err
	}

	// 签名直传。
	signURL, err = bucket.SignURL(key, oss.HTTPGet, 60)
	if err != nil {
		logger.Error("util GetSignUrl SignURL error:", err)
		return "", err
	}
	return signURL, err
}
