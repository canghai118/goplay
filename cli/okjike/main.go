package main

import (
	"github.com/gin-gonic/gin"
	"github.com/canghai118/goplay/cli/okjike/api"
)

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		s, err := api.SessionsCreate()
		if err != nil {
			c.JSON(500, gin.H{"error": err})
			return
		}
		// c.JSON(200, gin.H{"hello": s.UUID})
		png := api.CreateLoginQrCode(s)
		c.Writer.Header().Set("Content-Type", "image/png")
		c.Writer.Write(png)
	})

	r.Run(":8088")
}
