package model

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func MinioGetPng(path, contentId string) (imgURL string, err error) {
	key := path + "/" + contentId + ".png"

	req, _ := minioClient.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	// 署名付きURLの有効期限を設定
	imgURL, err = req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	log.Printf("Presigned URL: %s\n", imgURL)

	return imgURL, nil
}
