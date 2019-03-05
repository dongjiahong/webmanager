package task

type Task struct {
	TaskId      string     `json:"task_id"` // 任务id
	WorkerName  string     `json:"worker_name"`
	WorkerArgs  string     `json:"worker_args"`  // 下面func的参数
	WorkerFunc  WorkerFunc `json:"-" gorm:"-"`   // 这里的标签标示json和gorm都不序列化该字段
	ResultMsg   string     `json:"result_msg"`   // 结果信息，如：错误等
	ResultName  string     `json:"result_name"`  // 保存任务结果,输出的文件名
	ResultUrl   string     `json:"result_url"`   // 输出文件url
	ResultState int        `json:"result_state"` // 1: 完成,  2: 失败
	WorkerStart string     `json:"worker_start"` // 任务开始时间'20190203-02:03:21'
	WorkerEnd   string     `json:"worker_end"`   // 任务结束时间
}

// 参数： 任务名称和参数
// 返回： 任务结果和错误
type WorkerFunc func(string) (string, error)
