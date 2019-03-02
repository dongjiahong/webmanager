package main

import (
	"github.com/gin-gonic/gin"

	"webmanager/api"
	"webmanager/goque"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": name,
		})
	})
	r.GET("/edit/join", api.ApiJoin)

	r.GET("/dump/lq", api.ApiDumpLoopQueue)
	r.GET("/dump/dq", api.ApiDumpDoneQueue)

	goque.Init()

	r.Run()
}
