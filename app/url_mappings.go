package app

import (
	"github.com/dhanushhegde/bookstore_users-api/controllers"
)

func mapUrls() {
	router.GET("ping", controllers.Ping)
}
