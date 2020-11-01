package users

import (
	"strings"

	"github.com/dhanushhegde/bookstore_users-api/utils/errors"
)

// User represents a user
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// type User1 struct {
// 	Id          int64
// 	FirstName   string
// 	LastName    string
// 	Email       string
// 	DateCreated string
// }

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
