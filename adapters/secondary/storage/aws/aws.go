package storage

import (
	"bytes"
	"errors"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ecommerce/ports"
	"labix.org/v2/mgo/bson"
)

var (
	ErrorPutObject = errors.New("error save image")
)

type AWSStorage struct {
	session *session.Session
	bucket  string
}

func NewAWS(config aws.Config) ports.StorageService {
	s, _ := session.NewSession(&config)
	return &AWSStorage{
		session: s,
		bucket:  os.Getenv("AWS_BUCKET"),
	}
}

func (ss *AWSStorage) SaveFileAWS(file multipart.File, fileHeader string, fileSize int64, typeModel string) (*ports.SaveFileResponse, error) {
	buffer := make([]byte, fileSize)
	file.Read(buffer)
	fileName := typeModel + "/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader)
	imageURL := "https://" + ss.bucket + "." + "s3" + ".amazonaws.com/" + fileName

	_, err := s3.New(ss.session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(ss.bucket),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(fileSize)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return nil, ErrorPutObject
	}

	return &ports.SaveFileResponse{
		ImageKey: fileName,
		ImageURL: imageURL,
	}, nil
}

func (ss *AWSStorage) DeleteFileAWS(fileName string) (bool, error) {
	return false, nil
}
