package utility

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"labix.org/v2/mgo/bson"
)

//SendImageToAWS function
func SendImageToAWS(file multipart.File, fileHeader string, fileSize int64, typeModel string) (string, string) {
	secretKey := os.Getenv("AWS_SECRET_KEY")
	secretID := os.Getenv("AWS_SECRET_ID")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")

	buffer := make([]byte, fileSize)
	file.Read(buffer)

	sessionAWS := session.New(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(secretID, secretKey, ""),
	})

	fileName := typeModel + "/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader)
	_, err := s3.New(sessionAWS).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(fileName),
		ACL: aws.String("public-read"),
		Body: bytes.NewReader(buffer),
		ContentLength: aws.Int64(int64(fileSize)),
		ContentType: aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass: aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", ""
	}

	imageURL := "https://" + bucket + "." + "s3" + ".amazonaws.com/" + fileName

	return fileName, imageURL
}

//DeleteImageInAWS function
func DeleteImageInAWS(fileName string) bool {
	secretKey := os.Getenv("AWS_SECRET_KEY")
	secretID := os.Getenv("AWS_SECRET_ID")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET")
	
	svc := s3.New(session.New(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(secretID, secretKey, ""),
	}))

	input := &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &fileName,
	}

	_, err := svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
				return false
			}
		} else {
			fmt.Println(err.Error())
			return false
		}
	}

	return true
}