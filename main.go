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
	edit := r.Group("/edit")
	{
		edit.GET("/join", api.JoinMp4)
		edit.GET("/giftomp4", api.GifToMp4)
	}

	dump := r.Group("/dump")
	{
		dump.GET("/lq", api.ApiDumpLoopQueue)
		dump.GET("/dq", api.ApiDumpDoneQueue)
	}

	goque.Init()

	r.Run()
}
