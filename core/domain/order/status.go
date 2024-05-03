package order

import (
	"errors"
	"strings"
)

var (
	ErrorEmptyStatus = errors.New("empty status")
	ErrorWrongStatus = errors.New("wrong status")
)

func NewStatus(status string) (*string, error) {
	status = strings.TrimSpace(status)
	if len(status) == 0 {
		return nil, ErrorEmptyStatus
	}
	if status != "PENDENTE" && status != "CANCELADO" && status != "ENTREGUE" && status != "ANDAMENTO" {
		return nil, ErrorWrongStatus
	}

	return &status, nil
}
