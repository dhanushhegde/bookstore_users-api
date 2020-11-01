package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dhanushhegde/bookstore_users-api/domain/users"
	"github.com/dhanushhegde/bookstore_users-api/services"
	"github.com/dhanushhegde/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

//CreateUser function is used to create a user
func CreateUser(c *gin.Context) {
	var user users.User
	// var user1 users.User1

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// // bytes, _ := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(bytes))

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveError := services.CreateUser(user)

	if saveError != nil {
		//TODO: handle user creation error
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated, result)
	// fmt.Println(user)
	// fmt.Println(err)

	// c.String(http.StatusNotImplemented, "implement me!")
}

//GetUser function is used to get a user
func GetUser(c *gin.Context) {
	// var user User
	// if err:=services.
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

//SearchUser function is used to search for a user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
