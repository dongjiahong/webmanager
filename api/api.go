package api

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"

	"webmanager/ffmpeg"
	"webmanager/goque"
	"webmanager/model"
)

type Resp struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func newResp(msg string, data interface{}) *Resp {
	return &Resp{
		Message: msg,
		Data:    data,
	}
}

func (r *Resp) toJson() []byte {
	if r == nil {
		return nil
	}
	b, _ := json.Marshal(r)
	return b
}

// JoinMp4 拼接视频，需要的参数信息
// names: ["1.mp4", "2.mp4"] eg: curl "localhost:8080/edit/join?names=01.mp4,012.mp4"
func JoinMp4(c *gin.Context) {
	names := c.Query("names")

	goque.GetGoque().Add(ffmpeg.JoinVideo, names, "joinVideo")

	c.Data(200, "application/json", newResp("add task success", nil).toJson())
	return
}

// GifToMp4 将gif转为MP4
// curl "localhost:8080/edit/giftomp4?name=01.gif"
func GifToMp4(c *gin.Context) {
	name := c.Query("name")

	goque.GetGoque().Add(ffmpeg.GifToMp4, name, "gifToMp4")

	c.JSON(200, gin.H{
		"message": "add task ok",
	})
	return
}

func DumpLoopQueue(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": goque.GetGoque().Dump(),
	})
	return
}

func DumpTaskQueue(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	pageSize, err := strconv.Atoi(c.Param("pagesize"))
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	tasks, err := model.GetTasks(page, pageSize)
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	c.Data(200, "application/json", newResp("ok", tasks).toJson())
	return
}

// GetImg 获取image信息
// page: 第几页， pagesize：每页的数量
func GetImg(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	pageSize, err := strconv.Atoi(c.Param("pagesize"))
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	medias, err := model.GetMedia("img", page, pageSize)
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	c.Data(200, "application/json", newResp("ok", medias).toJson())
	return
}

func GetVideo(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	pageSize, err := strconv.Atoi(c.Param("pagesize"))
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	medias, err := model.GetMedia("video", page, pageSize)
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	c.Data(200, "application/json", newResp("ok", medias).toJson())
	return
}
