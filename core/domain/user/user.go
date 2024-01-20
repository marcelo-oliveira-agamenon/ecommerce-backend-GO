package user

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid"`
	Name       string
	Email      string
	Address    string
	ImageKey   string
	ImageURL   string
	Phone      string
	Password   string
	FacebookID string
	Birthday   string
	Gender     string
	Roles      pq.StringArray `gorm:"type:varchar(64)[]"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

var (
	ErrorUUID = errors.New("user id")
)

func NewUser(data User) (User, error) {
	id, errUUID := uuid.NewV4()
	if errUUID != nil {
		return User{}, ErrorUUID
	}

	return User{
		ID:         id,
		Name:       data.Name,
		Email:      data.Email,
		Address:    data.Address,
		Password:   data.Password,
		Phone:      data.Phone,
		FacebookID: data.FacebookID,
		ImageKey:   data.ImageKey,
		ImageURL:   data.ImageURL,
		Gender:     data.Gender,
		Roles:      data.Roles,
		Birthday:   data.Birthday,
	}, nil
}
