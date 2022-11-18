// Upload files individually
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Router() {
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST(&quot/upload&quot, func(c *gin.Context) {
		// this method using for upload single files
		file, _ := c.FormFile(&quotfile & quot)
		log.Println(file.Filename)
		var dst string = &quotdir / file / upload /
			c.SaveUploadedFile(file, dst)
	})
}
