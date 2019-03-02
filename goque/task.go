package goque

type Task struct {
	Id          string  `json:"id"`           // 任务id
	Result      string  `json:"result"`       // 保存任务结果
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
