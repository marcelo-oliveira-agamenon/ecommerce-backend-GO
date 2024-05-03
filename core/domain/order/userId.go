package order

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

var (
	ErrorEmptyUserId = errors.New("empty user id")
)

func NewUserId(userId string) (uuid.UUID, error) {
	userId = strings.TrimSpace(userId)
	if len(userId) == 0 {
		return uuid.UUID{}, ErrorEmptyUserId
	}

	nUserId := uuid.FromStringOrNil(userId)
	return nUserId, nil
}
