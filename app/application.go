package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//router is private variable

// var router = gin.Default()

func StartApplication() {
	mapUrls()
	// router.Run(addr ":8080")
	router.Run()
}
