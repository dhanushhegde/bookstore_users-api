package app

import (
	"github.com/dhanushhegde/bookstore_users-api/controllers/ping"

	"github.com/dhanushhegde/bookstore_users-api/controllers/user"
)

func mapUrls() {
	router.GET("ping", ping.Ping)
	router.GET("/users/:user_id", user.GetUser)
	// router.GET("/users/search", controllers.SearchUser)
	router.POST("/users", user.CreateUser)
}
