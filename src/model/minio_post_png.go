package model

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func MinioPostPng(path, contentId string, htmlPng []byte) error {
	key := path + "/" + contentId + ".png"

	// MinIOにファイルをアップロード
	_, err := minioClient.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(htmlPng),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return err
	}

	return nil
}
