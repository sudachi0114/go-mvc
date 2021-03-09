package models

import (
	"regexp"
	"time"

	"github.com/wcl48/valval"
)

type User struct {
	Id        int64
	Name      string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// User model のバリデーション
func UserValidate(user User) error {
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),                           // 名前は20文字以内
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`)), // 使えるのは英子文字とスペースだけ
		),
	})

	return Validator.Validate(user)
}
