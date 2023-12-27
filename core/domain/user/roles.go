package user

import (
	"errors"

	"github.com/lib/pq"
)

var (
	listRoles               = []string{"admin", "user"}
	ErrorWrongRoleOrInvalid = errors.New("invalid role added")
)

func NewRoles(roles pq.StringArray) (pq.StringArray, error) {
	valid := false
	for _, v := range listRoles {
		for _, v1 := range roles {
			if v1 == v {
				valid = true
			}
		}
	}
	if !valid {
		return nil, ErrorWrongRoleOrInvalid
	}

	return roles, nil
}

func IsRoleAdmin(roles pq.StringArray) bool {
	isAdmin := false
	for _, v := range roles {
		if v == "admin" {
			isAdmin = true
			break
		}
	}

	return isAdmin
}
