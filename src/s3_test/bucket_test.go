package s3_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"testing"
)

func TestBucketList(t *testing.T) {
	access_key := "3DNM7Z57L4UXGDPMG3FU"
	secret_key := "rWsCjB0u7GF4uPXTTaE0BU4rfNP33OE2WufJBJEt"
	end_point := "192.168.3.162:7480" //endpoint设置，不要动

	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(access_key, secret_key, ""),
		Endpoint:         aws.String(end_point),
		Region:           aws.String("default"),
		DisableSSL:       aws.Bool(true),
		//S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	})

	svc := s3.New(sess)

	input := &s3.ListBucketsInput{}
	result, err := svc.ListBuckets(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}