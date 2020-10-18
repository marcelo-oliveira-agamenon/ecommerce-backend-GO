package utility

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"labix.org/v2/mgo/bson"
)

//SendImageToAWS function
func SendImageToAWS(file multipart.File, fileHeader string, fileSize int64, typeModel string) (string) {
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
		return "Error upload image to AWS"
	}

	imageURL := "https://" + bucket + "." + "s3" + ".amazonaws.com/" + fileName

	return imageURL
}