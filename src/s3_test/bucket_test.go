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

var access_key = "3DNM7Z57L4UXGDPMG3FU"
var secret_key = "rWsCjB0u7GF4uPXTTaE0BU4rfNP33OE2WufJBJEt"
//var end_point = "192.168.3.162:7480"
var end_point = "object.yujiang.com:7480"
var bucket_name = "bucket1"
var object_name = "all.yml"
var object_version_id = "iGvhLxzNmMzN-3arUX-aG7ShepAUZB4"

func init() {

}

func svcErrProcess(err error)  {
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
	}
}

func svcGet() *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(access_key, secret_key, ""),
		Endpoint:         aws.String(end_point),
		Region:           aws.String("default"),
		DisableSSL:       aws.Bool(true),
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return s3.New(sess)
}

func TestPutBucketObjectLockExistingBuckets(t *testing.T) {
	svc := svcGet()
	input := &s3.PutObjectLockConfigurationInput{
		Bucket: aws.String(bucket_name),
	}
	resule, err := svc.PutObjectLockConfiguration(input)
	svcErrProcess(err)
	fmt.Println(resule)
}

func TestPutBucketObjectLockExistingBucketsWithObjectLockConfigurationEnable(t *testing.T) {
	/*
	<?xml version="1.0" encoding="UTF-8"?>
	<ObjectLockConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
	   <ObjectLockEnabled>string</ObjectLockEnabled>
	   <Rule>
	      <DefaultRetention>
	         <Days>integer</Days>
	         <Mode>string</Mode>
	         <Years>integer</Years>
	      </DefaultRetention>
	   </Rule>
	</ObjectLockConfiguration>
	*/
	svc := svcGet()
	input := &s3.PutObjectLockConfigurationInput{
		Bucket: aws.String(bucket_name),
		ObjectLockConfiguration:&s3.ObjectLockConfiguration{
			ObjectLockEnabled: aws.String("Enabled"),
			//ObjectLockEnabled: aws.String("Enabled"),
			Rule: &s3.ObjectLockRule{
				DefaultRetention: &s3.DefaultRetention{
					Mode: aws.String("GOVERNANCE"),
					//Mode: aws.String("COMPLIANCE"),
					/*
					$17 = {mode = {static npos = 18446744073709551615,
					    _M_dataplus = {<std::allocator<char>> = {<__gnu_cxx::new_allocator<char>> = {<No data fields>}, <No data fields>},
					      _M_p = 0x55c5c0952880 "COMPLIANCE"}, _M_string_length = 10, {_M_local_buf = "COMPLIANCE\000\067\320\177\000",
					      _M_allocated_capacity = 5638868800558157635}}, retain_until_date = {__d = {__r = 1619419423121442805}}}

					*/
					Days: aws.Int64(1),
				},
			},
		},
	}
	resule, err := svc.PutObjectLockConfiguration(input)
	svcErrProcess(err)
	fmt.Println(resule)
}

func TestPutBucketObjectLockExistingBucketsWithObjectLockConfigurationDisable(t *testing.T) {
	/*
		<?xml version="1.0" encoding="UTF-8"?>
		<ObjectLockConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
		   <ObjectLockEnabled>string</ObjectLockEnabled>
		   <Rule>
		      <DefaultRetention>
		         <Days>integer</Days>
		         <Mode>string</Mode>
		         <Years>integer</Years>
		      </DefaultRetention>
		   </Rule>
		</ObjectLockConfiguration>
	*/
	svc := svcGet()
	input := &s3.PutObjectLockConfigurationInput{
		Bucket: aws.String(bucket_name),
		ObjectLockConfiguration:&s3.ObjectLockConfiguration{
			ObjectLockEnabled: aws.String("Disable"),
			Rule: &s3.ObjectLockRule{
				DefaultRetention: &s3.DefaultRetention{
					Mode: aws.String("COMPLIANCE"),
					Days: aws.Int64(1),
				},
			},
		},
	}
	resule, err := svc.PutObjectLockConfiguration(input)
	svcErrProcess(err)
	fmt.Println(resule)
}

func TestGetBucketObjectLockEnabled(t *testing.T) {
	svc := svcGet()

	input := &s3.GetObjectLockConfigurationInput{
		Bucket: aws.String(bucket_name),
	}

	result, err := svc.GetObjectLockConfiguration(input)
	svcErrProcess(err)
	fmt.Println(result)
}

func TestBucketCreateEnableBucketObjectLock(t *testing.T) {
	svc := svcGet()

	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucket_name),
		ObjectLockEnabledForBucket: aws.Bool(true),
	}

	result, err := svc.CreateBucket(input)

	svcErrProcess(err)

	fmt.Println(result)
}


func TestBucketList(t *testing.T) {

	svc := svcGet()

	input := &s3.ListBucketsInput{}
	result, err := svc.ListBuckets(input)

	svcErrProcess(err)

	fmt.Println(result)
}

func TestDeleteObjectLockObject(t *testing.T) {
	svc := svcGet()
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket_name),
		Key: aws.String(object_name),
		BypassGovernanceRetention: aws.Bool(true),	// X_AMZ_BYPASS_GOVERNANCE_RETENTION
		VersionId: aws.String(object_version_id),
	}
	result, err := svc.DeleteObject(input)

	svcErrProcess(err)

	fmt.Println(result)
}

func TestListObjectVersion(t *testing.T) {
	svc := svcGet()
	input :=&s3.ListObjectVersionsInput{
		Bucket: aws.String(bucket_name),
		//KeyMarker: aws.String(object_name),
	}
	result, err := svc.ListObjectVersions(input)

	svcErrProcess(err)

	fmt.Println(result)
}