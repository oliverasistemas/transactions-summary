package s3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
)

import "os"

// ReadFileFromS3 reads the content of a file located in the specified S3 bucket.
// It takes the bucketName and key of the desired file as input parameters.
// It retrieves the object from S3 using the given bucketName and key and reads its content into a byte slice.
// The function returns the content of the file as a byte slice ([]byte) and any error that occurred during the process.
func ReadFileFromS3(bucketName, key string) ([]byte, error) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "sa-east-1"
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	rawObject, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing body")
		}
	}(rawObject.Body)

	buf := new(bytes.Buffer)

	_, err = buf.ReadFrom(rawObject.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
