package app

import (
	"github.com/dhanushhegde/bookstore_users-api/controllers/ping"

	"github.com/dhanushhegde/bookstore_users-api/controllers/user"
)

func mapUrls() {
	router.GET("ping", ping.Ping)
	// router.GET("/users/search", controllers.SearchUser)

	router.POST("/users", user.Create)
	router.GET("/users/:user_id", user.Get)
	router.PUT("/users/:user_id", user.Update)
	router.PATCH("/users/:user_id", user.Update)
	router.DELETE("/users/:user_id", user.Delete)

	router.GET("/internal/users/search", user.Search)
	router.GET("/users/login", user.Login)
}
