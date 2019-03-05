package goque

type Task struct {
	Id          string  `json:"id"`    // 任务id
	Start       string  `json:"start"` // 任务开始时间'20190203-02:03:21'
	End         string  `json:"end"`   // 任务结束时间
	Result      string  `json:"result"`
	ResultName  string  `json:"result_name"`  // 保存任务结果,输出的文件名
	RusultUrl   string  `json:"result_url"`   // 输出文件url
	ResultState int     `json:"result_state"` // 1: 完成,  2: 等待, 3: 失败
	Worker      *Worker `json:"worker"`
}

// 参数： 任务名称和参数
// 返回： 任务结果和错误
type workerFunc func(string) (string, error)

type Worker struct {
	Name string `json:"name"`
	wf   workerFunc
	args string `json:"args"`
}
