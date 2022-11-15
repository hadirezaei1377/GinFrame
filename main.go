// Using GET, POST, PUT, PATCH, DELETE and OPTIONS
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// In this case, by sending a string to the address above (user/:name/), which replaces our name, which is as follows
	// /user/:name
	// /user/Hadi

	// Query String can be one or more parameters
	// the address that we should have is "welcome?firstname=Jane&lastname=Doe"
	router.GET("/welcome", func(c *gin.Context) {
		// if client doesnt send anything show this by default
		firstname := c.DefaultQuery("firstname", "it was empty!")
		// change request to querystring
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

	})
}
