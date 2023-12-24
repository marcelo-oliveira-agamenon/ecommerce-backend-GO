package user

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	ErrorInvalidBirthday = errors.New("invalid birthday")
)

func NewBirthday(birth string) (string, error) {
	if match, err := regexp.MatchString(`\d{1,2}\/\d{1,2}\/\d{2,4}`, birth); err != nil || !match {
		return "", ErrorInvalidBirthday
	}

	ty := time.Now().Year()
	spBi := strings.Split(birth, "/")
	by, _ := strconv.Atoi(spBi[2])
	if ty - by > 150 {
		return "", ErrorInvalidBirthday
	}

	return birth, nil
}