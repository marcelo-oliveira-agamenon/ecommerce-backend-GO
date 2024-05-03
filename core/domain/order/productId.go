package order

import (
	"errors"
	"strings"

	"github.com/lib/pq"
)

var (
	ErrorEmptyProductId = errors.New("empty product id")
)

func NewProductId(prodId string) (pq.StringArray, error) {
	prodId = strings.TrimSpace(prodId)
	if len(prodId) == 0 {
		return nil, ErrorEmptyProductId
	}

	prSp := strings.Split(prodId, ",")
	return prSp, nil
}
