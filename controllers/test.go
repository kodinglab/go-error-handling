package controllers

import (
	"error-handling/middleware"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestController struct
type TestController struct {
}

// GetNormalRoute method
func (ctr *TestController) GetNormalRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"mesage": "This is valid route",
	})
}

// GetErrorRoute method
func (ctr *TestController) GetErrorRoute(c *gin.Context) {
	err := errors.New("This is error route")
	if err != nil {
		middleware.Log{}.Error(err)

		// and you can display output error (500 status code)
		middleware.ErrorOutput(c, err)
	}
}

// GetAnotherErrorRoute method
func (ctr *TestController) GetAnotherErrorRoute(c *gin.Context) {
	err := errors.New("This is another error route")
	if err != nil {
		middleware.Log{}.Error(err)

		// and also you can display output error (500 status code)
		middleware.ErrorOutput(c, err)
	}
}
