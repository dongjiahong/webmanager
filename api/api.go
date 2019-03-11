package api

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"webmanager/ffmpeg"
	"webmanager/goque"
	"webmanager/model"
	"webmanager/util"
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
	mp, err := newMaybeParams(c)
	if err != nil {
		log.Println("[JoinMp4] get params err: ", err.Error())
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	goque.GetGoque().Add(ffmpeg.JoinVideo, mp.FileNames, "joinVideo")

	c.Data(200, "application/json", newResp("add task success", nil).toJson())
	return
}

// GifToMp4 将gif转为MP4
// curl "localhost:8080/edit/giftomp4?name=01.gif"
func GifToMp4(c *gin.Context) {
	mp, err := newMaybeParams(c)
	if err != nil {
		log.Println("[GifToMp4] get params err: ", err.Error())
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	goque.GetGoque().Add(ffmpeg.GifToMp4, mp.FileName, "gifToMp4")

	c.Data(200, "application/json", newResp("add task success", nil).toJson())
	return
}

func GetTaskQueue(c *gin.Context) {
	mp, err := newMaybeParams(c)
	if err != nil {
		log.Println("[GetTaskQueue] get params err: ", err.Error())
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	tasks, err := model.GetTasks(mp.Page, mp.PageSize)
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
	mp, err := newMaybeParams(c)
	if err != nil {
		log.Println("[GetImg] get params err: ", err.Error())
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	medias, err := model.GetMedia("img", mp.Page, mp.PageSize)
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	c.Data(200, "application/json", newResp("ok", medias).toJson())
	return
}

func GetVideo(c *gin.Context) {
	mp, err := newMaybeParams(c)
	if err != nil {
		log.Println("[GetVideo] get params err: ", err.Error())
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	medias, err := model.GetMedia("video", mp.Page, mp.PageSize)
	if err != nil {
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}
	c.Data(200, "application/json", newResp("ok", medias).toJson())
	return
}

func UploadMedia(c *gin.Context) {
	var mp model.Media
	if err := c.ShouldBindJSON(&mp); err != nil {
		log.Println("[UploadMedia] decode json err: ", err.Error(), " req: ", *c.Request)
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	if mp.Source == "upload" { // 我们要去下载该视频
		kind := func() string {
			if mp.MediaType == "mp4" {
				return "video"
			}
			return "img"
		}()
		fileName := func() string {
			if strings.HasSuffix(mp.FileName, "mp4") {
				return mp.FileName
			}
			return mp.FileName + ".mp4"
		}()
		filePath := util.GetCommonPath(kind) + fileName
		if err := util.DownloadWithUrl(mp.Url, filePath); err != nil {
			log.Println("[UploadMedia] download file err: ", err.Error())
			c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
			return
		}
	}

	if err := model.WriteMediaToDB(&mp); err != nil {
		log.Println("[UploadMedia] write db err: ", err.Error())
		c.Data(200, "application/json", newResp(err.Error(), nil).toJson())
		return
	}

	c.Data(200, "application/json", newResp("ok", nil).toJson())
	return
}
