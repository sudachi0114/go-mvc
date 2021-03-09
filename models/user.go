package models

import (
	"time"
)

type User struct {
	Id        int64
	Name      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
