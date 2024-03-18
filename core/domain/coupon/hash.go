package coupon

import (
	"errors"
	"strings"
)

var (
	ErrorHashTooShort = errors.New("title too short")
	ErrorEmptyHash    = errors.New("empty hash")
)

func NewHash(hash string) (*string, error) {
	hash = strings.TrimSpace(hash)
	if len(hash) == 0 {
		return nil, ErrorEmptyHash
	}
	if len(hash) < 6 {
		return nil, ErrorHashTooShort
	}

	return &hash, nil
}
