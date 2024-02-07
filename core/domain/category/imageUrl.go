package category

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrorEmptyimageUrl   = errors.New("empty image url")
	ErrorInvalidimageUrl = errors.New("invalid image url")
)

func NewImageUrl(imageUrl string) (*string, error) {
	imageUrl = strings.TrimSpace(imageUrl)
	bucket := os.Getenv("AWS_BUCKET")
	if len(imageUrl) == 0 {
		return nil, ErrorEmptyName
	}
	if strings.Contains(imageUrl, bucket) {
		return nil, ErrorInvalidimageUrl
	}

	return &imageUrl, nil
}
