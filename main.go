package main

import (
	"github.com/dongjiahong/gotools"
	"github.com/gin-gonic/gin"

	"webmanager/api"
	"webmanager/goque"
	"webmanager/model"
	"webmanager/util"
)

type Conf struct {
	ModelConf *model.Conf `json:"model"`
	UtilConf  *util.Conf  `json:"util"`
}

func main() {
	r := gin.Default()

	get := r.Group("/get")
	{
		get.GET("/img/:page/:pagesize", api.GetImg)
		get.GET("/video/:page/:pagesize", api.GetVideo)
		get.GET("/tq/:page/:pagesize", api.GetTaskQueue)
	}

	edit := r.Group("/edit")
	{
		edit.GET("/join", api.JoinMp4)
		edit.GET("/giftomp4", api.GifToMp4)
	}

	r.POST("/upload", api.UploadMedia)

	media := r.Group("/media")
	{
		media.Static("/img", "./media/img")
		media.Static("/video", "./media/video")
	}

	r.StaticFile("/", "./dist/index.html")
	s := r.Group("/static")
	{
		s.Static("/css", "./dist/static/css")
		s.Static("/fonts", "./dist/static/fonts")
		s.Static("/js", "./dist/static/js")
	}

	var conf Conf
	if err := gotools.DecodeJsonFromFile("config/manager.json", &conf); err != nil {
		panic(err)
	}
	goque.Init()
	util.Init(conf.UtilConf)

	if err := model.NewDb(conf.ModelConf); err != nil {
		panic(err)
	}
	defer model.Close()

	r.Run()
}
