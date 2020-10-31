package users

import{

	"github.com/dhanushhegde/bookstore_users-api/utils/errors"
}

var {
	usersDB = make(map[int64]*User)
}

func (user User) Get() (*User, *errors.RestErr) {
	result := userDB[user.Id]
	if result == nil{
		return nil,errors.NewNotFoundError{fmt.Sprintf{"user %d not found",user.Id}} 
	}
	return nil, nil
}
func (user User) Save() *errors.RestErr {
	return nil
}
