package main

import (
	"github.com/gin-gonic/gin"
)

func FileReader() {
	// We have a form called form.html.
	// and a text box named message.
	// We are going to send the information from this form to the server.
	router.POST(&quot/form_post&quot, func(c *gin.Context) {
		// c.PostForm is main command for this task
		message := c.PostForm(&quotmessage & quot)
	})

}
