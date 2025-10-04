package s3wrapper

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/maooz4426/aws-sdk-go-playground/config"
)

type S3API interface {
	ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error)
}
type S3Client struct {
	Client S3API
}

func NewS3Client() (*S3Client, error) {
	sdkConfig := config.NewConfig()
	s3Client := s3.NewFromConfig(*sdkConfig.AWSConfig)
	return &S3Client{s3Client}, nil
}

func (s S3Client) ListBuckets() (*s3.ListBucketsOutput, error) {
	result, err := s.Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	return result, nil
}
