package domain

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"stori/pkg/aws/s3"
)

const (
	envS3BucketName  = "S3_BUCKET_NAME"
	envS3FileKey     = "S3_FILE_KEY"
	localCSVFilePath = "LOCAL_CSV_FILE_PATH"
)

// extractTransactions reads transaction data from a CSV file or an S3 bucket,
// depending on the presence of the LOCAL_CSV_FILE_PATH environment variable.
// If LOCAL_CSV_FILE_PATH is set, it reads the CSV data from the local file.
// Otherwise, it fetches the CSV data from the specified S3 bucket and key using
// the environment variables S3_BUCKET_NAME and S3_FILE_KEY.
// The function returns a slice of Transaction objects and an error if any.
func extractTransactions() ([]Transaction, error) {
	localFile := os.Getenv(localCSVFilePath)
	var reader *csv.Reader

	if localFile != "" {
		fileContent, err := os.Open(localFile)
		if err != nil {
			return nil, err
		}
		defer fileContent.Close()
		reader = csv.NewReader(fileContent)

	} else {
		bucket, key := getS3Env()
		bytesSlice, err := s3.ReadFileFromS3(bucket, key)
		if err != nil {
			return nil, err
		}
		reader = csv.NewReader(bytes.NewReader(bytesSlice))
	}

	return processReader(reader)
}

func getS3Env() (bucket, key string) {
	bucket = os.Getenv(envS3BucketName)
	if bucket == "" {
		log.Fatalf("please, set %s environment variable", envS3BucketName)
	}

	key = os.Getenv(envS3FileKey)
	if bucket == "" {
		log.Fatalf("please, set %s environment variable", envS3FileKey)
	}

	return bucket, key
}
