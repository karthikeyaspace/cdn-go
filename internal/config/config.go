package config

//https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type Config struct {
	AWSRegion  string
	AccessKey  string
	SecretKey  string
	BucketName string
	Port       string
}

var (
	cfg      *Config
	s3Client *s3.Client
	once     sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		}
		cfg = &Config{
			AWSRegion: os.Getenv("AWS_REGION"),
			AccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
			SecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
			BucketName: os.Getenv("AWS_BUCKET_NAME"),
			Port: os.Getenv("PORT"),
		}

	})

	return cfg
}

func initS3Client() {
	awsCfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cfg.AWSRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	s3Client = s3.NewFromConfig(awsCfg)
}

func GetS3Client() *s3.Client {
	if s3Client == nil {
		initS3Client()
	}
	return s3Client
}