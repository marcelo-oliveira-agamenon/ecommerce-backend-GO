package coupon

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrorEmptyExpDate = errors.New("empty expiration date")
	ErrorParseTime    = errors.New("in expiration date, check format")
)

func NewExpireDate(expDate string) (*time.Time, error) {
	expDate = strings.TrimSpace(expDate)
	if len(expDate) == 0 {
		return nil, ErrorEmptyExpDate
	}

	expTime, errT := time.Parse("2006-01-02T15:04:05", expDate)
	if errT != nil {
		return nil, ErrorParseTime
	}

	return &expTime, nil
}
