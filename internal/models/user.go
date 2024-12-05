package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId      uuid.UUID `json:"Id"`
	UserName    string    `json:"UserName"`
	Password    string    `json:"Password"`
	CreatedDate time.Time `json:"CreatedDate"`
	IsAdmin     bool      `json:"IsAdmin"`
}
