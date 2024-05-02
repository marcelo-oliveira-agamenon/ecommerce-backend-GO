package logs

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyType = errors.New("empty type")
	ErrorWrongType = errors.New("wrong type")
	Types          = []string{"user", "product", "category", "coupon", "favorite", "order", "payment", "admin"}
)

func NewType(typ string) (*string, error) {
	typ = strings.TrimSpace(typ)
	if len(typ) == 0 {
		return nil, ErrorEmptyType
	}

	mtc := 0
	for _, v := range Types {
		if v == typ {
			mtc++
		}
	}

	if mtc != 1 {
		return nil, ErrorWrongType
	}

	return &typ, nil
}
