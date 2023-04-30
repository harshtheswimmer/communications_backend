package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestController() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := gin.H{
			"Code": 0,
			"data": "We have alot of data",
			"msg":  "Welcome You"}
		c.JSON(http.StatusOK, data)
	}
}
