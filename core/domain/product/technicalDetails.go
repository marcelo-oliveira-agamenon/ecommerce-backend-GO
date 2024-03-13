package product

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyTechnical              = errors.New("empty technical description")
	ErrorTechnicalDescriptionTooLong = errors.New("technical description too long")
)

func NewTechnicalDescription(technicalDescription string) (*string, error) {
	technicalDescription = strings.TrimSpace(technicalDescription)
	if len(technicalDescription) == 0 {
		return nil, ErrorEmptyTechnical
	}
	if len(technicalDescription) > 5000 {
		return nil, ErrorTechnicalDescriptionTooLong
	}

	return &technicalDescription, nil
}
