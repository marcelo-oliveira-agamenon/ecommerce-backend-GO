package logs

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyDescription    = errors.New("empty description")
	ErrorDescriptionTooShort = errors.New("description too short")
	ErrorDescriptionTooLong  = errors.New("description too long")
)

func NewDescription(des string) (*string, error) {
	des = strings.TrimSpace(des)
	if len(des) == 0 {
		return nil, ErrorEmptyDescription
	}

	if len(des) < 10 {
		return nil, ErrorDescriptionTooShort
	}
	if len(des) > 255 {
		return nil, ErrorDescriptionTooLong
	}

	return &des, nil
}
