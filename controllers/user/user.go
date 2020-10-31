package user

import (
	"net/http"

	"github.com/dhanushhegde/bookstore_users-api/domain/users"
	"github.com/dhanushhegde/bookstore_users-api/services"
	"github.com/dhanushhegde/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

//CreateUser function is used to create a user
func CreateUser(c *gin.Context) {
	var user users.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {

		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveError := services.CreateUser(user)

	if saveError != nil {
		//TODO: handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
	// fmt.Println(user)
	// fmt.Println(err)

	// c.String(http.StatusNotImplemented, "implement me!")
}

//GetUser function is used to get a user
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

//SearchUser function is used to search for a user
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
