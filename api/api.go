package api

import (
	"github.com/gin-gonic/gin"

	"webmanager/ffmpeg"
	"webmanager/goque"
)

// 拼接视频，需要的参数信息
// names: ["1.mp4", "2.mp4"]
func ApiJoin(c *gin.Context) {
	names := c.Query("names")

	goque.GetGoque().Add(ffmpeg.JoinVideo, names)

	c.JSON(200, gin.H{
		"message": "add task ok",
	})
}

// 将gif转为MP4
func ApiGifToMp4(c *gin.Context) {

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
