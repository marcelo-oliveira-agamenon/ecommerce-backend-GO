package user

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	ErrorEmptyBirthday      = errors.New("empty birthday")
	ErrorInvalidBirthday    = errors.New("invalid birthday")
	ErrorBirthdayOutOfRange = errors.New("birthday date out of range")
)

func NewBirthday(birth string) (string, error) {
	birth = strings.TrimSpace(birth)
	if len(birth) == 0 {
		return "", ErrorEmptyBirthday
	}

	if match, err := regexp.MatchString(`\d{1,2}\/\d{1,2}\/\d{2,4}`, birth); err != nil || !match {
		return "", ErrorInvalidBirthday
	}

	ty := time.Now().Year()
	spBi := strings.Split(birth, "/")
	by, _ := strconv.Atoi(spBi[2])
	if ty-by > 150 {
		return "", ErrorBirthdayOutOfRange
	}

	return birth, nil
}
