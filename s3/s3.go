package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"os"
)

type s3client struct {
	s3     *s3.S3
	bucket string
}

func NewS3Client(endpoint string, bucket string, key string, secret string, region string) *s3client {
	s, err := session.NewSession(&aws.Config{
		Endpoint:    aws.String(endpoint),
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
	})

	if err != nil {
		fmt.Printf("Unable create New S3 s %v\n", err)
		os.Exit(1)
	}

	return &s3client{s3.New(s), bucket}
}

func (c *s3client) GetContent(path string) (io.Reader, error) {
	resp, err := c.getContent(path)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (c *s3client) getContent(path string) (resp *s3.GetObjectOutput, err error) {
	resp, err = c.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(path),
	})

	return
}

func (c *s3client) headContent(path string) (resp *s3.HeadObjectOutput, err error) {
	resp, err = c.s3.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(path),
	})

	return
}
