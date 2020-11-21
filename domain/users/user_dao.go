package users

import (
	"fmt"

	"github.com/dhanushhegde/bookstore_users-api/datasource/mysql/users_db"
	"github.com/dhanushhegde/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
	querySelect     = "SELECT*FROM users"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	if err := users_db.Client.Ping(); err != nil {
		panic(err)

	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
func (user *User) Save() *errors.RestErr {
	// stmt, _ := users_db.Client.Prepare("SELECT*FROM users")
	// users_db.CLient.
	userDbTemp := users_db.Client
	if userDbTemp == nil {
		fmt.Println("Db null")
	}
	results, err := userDbTemp.Exec("SELECT*FROM users")
	// results, err := users_db.Client.Exec("SELECT*FROM users")
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	fmt.Println(results)
	// stmt, err := users_db.Client.Prepare(queryInsertUser)
	// if err != nil {
	// 	return errors.NewInternalServerError(err.Error())
	// }
	// defer stmt.Close()
	// insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	// if err != nil {
	// 	return errors.NewInternalServerError(
	// 		fmt.Sprintf("error while trying to save user: %s", err.Error()))
	// }
	// userId, err := insertResult.LastInsertId()
	// if err != nil {
	// 	return errors.NewInternalServerError(
	// 		fmt.Sprintf("error while trying to save user: %s", err.Error()))
	// }
	// user.Id = userId
	user.Id = 1
	return nil
}
