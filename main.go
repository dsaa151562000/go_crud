package main

import (
	"github.com/dsaa151562000/go_crud/db"
	"github.com/dsaa151562000/go_crud/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})
	db.Init()
	server.Init()
	db.Close()
}
