package main

import (
	"error-handling/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// handler for if no route available
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Route Not Found"})
	})

	// create test controller
	testCtr := &controllers.TestController{}

	restgo := router.Group("error-handling")
	{
		restgo.GET("/test", testCtr.GetNormalRoute)
		restgo.GET("/test/error", testCtr.GetErrorRoute)
		restgo.GET("/test/another-error", testCtr.GetAnotherErrorRoute)
	}

	router.Run(":8080")
}
