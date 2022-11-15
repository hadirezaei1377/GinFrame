// Using GET, POST, PUT, PATCH, DELETE and OPTIONS
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)
	// default port (8080)
	router.Run()

	// pass parameters by URL
}
