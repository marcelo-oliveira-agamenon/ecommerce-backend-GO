package product

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyDescription   = errors.New("empty description")
	ErrorDescriptionTooLong = errors.New("description too long")
)

func NewDescription(description string) (*string, error) {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		return nil, ErrorEmptyDescription
	}
	if len(description) > 1000 {
		return nil, ErrorDescriptionTooLong
	}

	return &description, nil
}
