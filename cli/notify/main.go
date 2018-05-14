package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.Writer.WriteString("/test")
	})

	router.Run()

	http.Handle("/bb", router)

	http.ListenAndServe(":8881", nil)
}

