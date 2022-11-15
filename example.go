package main

import "github.com/gin-gonic/gin"

func testing() {
	// using server
	r := gin.Default()
	// testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
}
