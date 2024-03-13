package favorite

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

var (
	ErrorIncorrectUserId = errors.New("incorrect user id")
)

func NewUserId(userId uuid.UUID) (*uuid.UUID, error) {
	if len(userId.String()) != 36 {
		return nil, ErrorIncorrectUserId
	}
	if !strings.Contains(userId.String(), "-") {
		return nil, ErrorIncorrectUserId
	}

	return &userId, nil
}
