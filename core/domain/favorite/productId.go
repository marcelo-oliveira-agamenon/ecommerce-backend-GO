package favorite

import (
	"errors"
	"strings"
)

var (
	ErrorIncorrectProduct = errors.New("incorrect product id")
	ErrorEmptyProducid    = errors.New("empty product id")
)

func NewProductId(prodId string) (*string, error) {
	prodId = strings.TrimSpace(prodId)
	if len(prodId) == 0 {
		return nil, ErrorEmptyProducid
	}
	if len(prodId) < 3 {
		return nil, ErrorIncorrectProduct
	}

	return &prodId, nil
}
