# webmanager

该服务主要是用来使用FFmpeg处理视频，提供前端工具

## 使用 go modules 管理，go 版本为1.12

```sh
# 素材表
CREATE TABLE media(
    id INT PRIMARY KEY AUTO_INCREMENT comment '自增主键',
    name VARCHAR(256) NOT NULL comment '文件名',
    title TEXT comment '素材title',
    description TEXT comment '素材描述',
    video_or_pic INT NOT NULL comment '是视频还是图片，1:video, 2:pic',
    media_type VARCHAR(128) NOT NULL comment '素材类型，mp4,gif等',
    media_tag VARCHAR(256) comment '素材标签，搞笑、热门等',
    url TEXT NOT NULL comment '素材链接',
    ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY `idx_id` (`id`),
    KEY `idx_name` (`name`)
) ENGINE = INNODB DEFAULT CHARSET = utf8
# 插入一条新数据
INSERT INTO media ( name, title, description, video_or_pic,media_type, media_tag, url) VALUES ('02.mp4', '励志小视频', '小孩视频真励志', 'video','mp4', '搞笑,励志', 'http://localhost:8080/media/video/02.mp4');

# 任务表
CREATE TABLE tasks(
    id INT PRIMARY KEY AUTO_INCREMENT comment '自增主键',
    task_id VARCHAR(256) NOT NULL comment '任务id',
    worker_name VARCHAR(256) NOT NULL comment '任务名',
    worker_args TEXT NOT NULL comment '任务参数',
    result_msg TEXT comment '任务结果信息',
    result_name TEXT comment '任务生产文件名',
    result_url TEXT comment '任务生产文件url',
    result_state INT comment '任务状态：1：完成，2：等待，3：失败',
    START VARCHAR(128) comment '任务开始时间',
    END VARCHAR(128) comment '任务结束时间',
    KEY `idx_id` (`id`),
    KEY `idx_task_id` (`task_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8
```
