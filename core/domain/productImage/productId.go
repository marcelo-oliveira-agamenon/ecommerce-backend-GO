package productImage

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyProductId = errors.New("empty product id")
)

func NewProductId(prodId string) (*string, error) {
	prodId = strings.TrimSpace(prodId)
	if len(prodId) == 0 {
		return nil, ErrorEmptyProductId
	}

	return &prodId, nil
}
