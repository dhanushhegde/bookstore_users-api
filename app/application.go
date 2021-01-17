package app

import (
	"github.com/dhanushhegde/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//router is private variable

// var router = gin.Default()

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application...")
	router.Run()
}
