package model

import (
	"regexp"

	common "github.com/minhquy1903/flashcard-api/internal/common/model"
)

type User struct {
	common.BaseModel
	Email    string
	Name     string
	Password string
}

func (d User) IsEmailValid() bool {
	// regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// match email address against pattern
	match, err := regexp.MatchString(pattern, d.Email)
	if err != nil {
		return false
	}

	return match
}

func (d User) IsNotEmpty() bool {
	if d.Email == "" || d.Password == "" {
		return false
	}

	return true
}

func (d User) TableName() string {
	return "user"
}
