package model

import (
	"webmanager/task"
)

type Media struct {
	Id          int    `json:"id"`           // 素材的唯一id
	Name        string `json:"name"`         // 素材的名字
	Title       string `json:"title"`        // 素材的title
	Description string `json:"description"`  // 素材的描述
	VideoOrPic  string `json:"video_or_pic"` // 素材是视频还是图片，video,picture
	MediaType   string `json:"media_type"`   // 如：mp4, gif
	MediaTag    string `json:"media_tag"`    // 素材的关键字，如："搞笑", "流行", 多个tag用逗号分隔
	Url         string `json:"url"`          // 素材的url
	Ts          string `json:"ts"`           // 时间戳
}

func GetMedia(videoOrPic string, page, pageSize int) ([]Media, error) {
	medias := make([]Media, 0, 10)
	err := gdb.Offset(page*pageSize).Limit(pageSize).Where("video_or_pic = ?", videoOrPic).Find(&medias).GetErrors()
	if len(err) != 0 {
		return nil, err[0]
	}

	return medias, nil
}

func GetTasks(page, pageSize int) ([]task.Task, error) {
	tasks := make([]task.Task, 0, 10)
	err := gdb.Offset(page * pageSize).Limit(pageSize).Find(&tasks).GetErrors()
	if len(err) != 0 {
		return nil, err[0]
	}

	return tasks, nil
}

func WriteTaskToDB(t *task.Task) error {
	gdb.NewRecord(t)
	errs := gdb.Create(t).GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}
