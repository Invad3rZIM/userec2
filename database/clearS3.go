package database

import (
	"fmt"
	"hypeman-userec2/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func (h *Database) ClearS3() error {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(constants.S3_REGION)})

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.

	if err != nil {
		return err
	}
	svc := s3.New(sess)

	input := &s3.ListObjectsInput{
		Bucket:  aws.String(constants.S3_BUCKET),
		MaxKeys: aws.Int64(1000),
	}

	result, err := svc.ListObjects(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
		}
	}

	for _, y := range result.Contents {
		svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(constants.S3_BUCKET), Key: y.Key})
	}

	return nil

}
