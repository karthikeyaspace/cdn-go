package service

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/karthikeyaspace/cdn-go/internal/config"
)

func UploadToS3(key string, file io.Reader) error {
	s3Client := config.GetS3Client()
	cfg := config.LoadConfig()
	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(cfg.BucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String("text/html"),
	})

	return err
}

func GetFilesFromS3(key string) ([]byte, error) {
	s3Client := config.GetS3Client()
	cfg := config.LoadConfig()

	file, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(cfg.BucketName),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}
	defer file.Body.Close()

	fileContent, err := io.ReadAll(file.Body)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
