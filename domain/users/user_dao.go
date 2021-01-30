package users

import (
	"fmt"

	"github.com/dhanushhegde/bookstore_users-api/datasource/mysql/users_db"
	"github.com/dhanushhegde/bookstore_users-api/logger"
	"github.com/dhanushhegde/bookstore_users-api/utils/errors"
	"github.com/dhanushhegde/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name,last_name,email,date_created,status,password) VALUES(?,?,?,?,?,?);"
	queryGetUser                = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?,last_name=?,email=? WHERE Id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus       = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE status=?;"
	errorNoRows                 = "no rows in result set"
	indexUniqueEmail            = "email_UNIQUE"
	queryFindByEmailAndPassword = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE email=? and password=?"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error while trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {

		// return mysql_utils.ParseError(getErr)

		logger.Error("error while trying to get user by id", getErr)
		return errors.NewInternalServerError("database error")

	}

	return nil
}
func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error while trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")

	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		// return errors.NewInternalServerError(fmt.Sprintf("error while trying to save user: %s", err.Error()))
		logger.Error("error while trying to save user", err)
		return errors.NewInternalServerError("database error")
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		// return errors.NewInternalServerError(err.Error())
		logger.Error("error while trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)

	if err != nil {
		// return mysql_utils.ParseError(err)
		logger.Error("error while trying to update user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil

}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		// return errors.NewInternalServerError(err.Error())
		logger.Error("error while trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id)
	if err != nil {
		// return mysql_utils.ParseError(err)
		logger.Error("error while trying to delete user", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		// return nil, errors.NewInternalServerError(err.Error())
		logger.Error("error while trying to prepare find user statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
		logger.Error("error while trying to find user", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			// return nil, mysql_utils.ParseError(err)
			logger.Error("error while trying to scan user into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		// return nil, errors.NewInternalServerError(err.Error())
		logger.Error("error while trying to prepare find user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email, user.Password)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {

		// return mysql_utils.ParseError(getErr)

		logger.Error("error while trying to get user email and password", getErr)
		return errors.NewInternalServerError("database error")

	}

	return nil

}
