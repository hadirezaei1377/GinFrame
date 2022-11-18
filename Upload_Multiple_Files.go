// Upload multiple files
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Router() {
	router.POST(&quot/upload&quot, func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File[&quotupload[]&quot]
		for _, file := range files {
		c.SaveUploadedFile(file, dst)
		}
		})
	})

