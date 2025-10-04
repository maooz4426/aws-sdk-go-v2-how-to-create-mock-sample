package main

import (
	"fmt"

	"github.com/joho/godotenv"
	s3client "github.com/maooz4426/aws-sdk-go-playground/s3_wrapper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	fmt.Println("start")

	s3client, err := s3client.NewS3Client()
	if err != nil {
		panic(err)
	}
	result, err := s3client.ListBuckets()
	if err != nil {
		panic(err)
	}

	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name)
	}

}
