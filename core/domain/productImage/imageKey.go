package productImage

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyName       = errors.New("empty name")
	ErrorEmptyImageKey   = errors.New("empty image key")
	ErrorImageKeyTooLong = errors.New("image key too long")
)

func NewImageKey(imageKey string) (*string, error) {
	imageKey = strings.TrimSpace(imageKey)
	if len(imageKey) == 0 {
		return nil, ErrorEmptyName
	}
	if len(imageKey) > 5000 {
		return nil, ErrorImageKeyTooLong
	}

	return &imageKey, nil
}
