package api

import (
	"github.com/gin-gonic/gin"

	"webmanager/ffmpeg"
	"webmanager/goque"
)

// JoinMp4 拼接视频，需要的参数信息
// names: ["1.mp4", "2.mp4"] eg: curl "localhost:8080/edit/join?names=01.mp4,012.mp4"
func JoinMp4(c *gin.Context) {
	names := c.Query("names")

	goque.GetGoque().Add(ffmpeg.JoinVideo, names)

	c.JSON(200, gin.H{
		"message": "add task ok",
	})
}

// GifToMp4 将gif转为MP4
// curl "localhost:8080/edit/giftomp4?name=01.gif"
func GifToMp4(c *gin.Context) {
	name := c.Query("name")

	goque.GetGoque().Add(ffmpeg.GifToMp4, name)

	c.JSON(200, gin.H{
		"message": "add task ok",
	})
}

func ApiDumpLoopQueue(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": goque.GetGoque().Dump(),
	})
}

func ApiDumpDoneQueue(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": goque.GetGoque().DumpDone(),
	})
}
