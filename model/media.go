package model

import (
	"fmt"

	"webmanager/task"
)

type Media struct {
	Id          int    `json:"id" gorm:"-"`  // 素材的唯一id
	FileName    string `json:"file_name"`    // 素材的名字
	Title       string `json:"title"`        // 素材的title
	Description string `json:"description"`  // 素材的描述
	UseNum      int    `json:"use_num"`      // 素材被使用的次数
	Source      string `json:"source"`       // 素材来源: spider:爬虫, upload:上传, job:任务合成
	VideoOrPic  string `json:"video_or_pic"` // 素材是视频还是图片，video,picture
	MediaType   string `json:"media_type"`   // 如：mp4, gif
	MediaTag    string `json:"media_tag"`    // 素材的关键字，如："搞笑", "流行", 多个tag用逗号分隔
	Url         string `json:"url"`          // 素材的url
	Ts          string `json:"ts" gorm:"-"`  // 时间戳
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
	if t == nil {
		return fmt.Errorf("wirte task find nil")
	}
	gdb.NewRecord(t)
	errs := gdb.Create(t).GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}

func WriteMediaToDB(m *Media) error {
	if m == nil {
		return fmt.Errorf("wirte media find nil")
	}
	gdb.NewRecord(m)
	errs := gdb.Create(m).GetErrors()
	if len(errs) > 0 {
		return errs[0]
	}
	return nil
}
