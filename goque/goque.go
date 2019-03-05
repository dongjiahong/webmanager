package goque

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"webmanager/model"
	"webmanager/task"
)

type Goque struct {
	lq *Queue `json:"loop_queue"` // 任务接受队列 loop queue
	tk *time.Ticker
	wg sync.WaitGroup
}

var gGoque *Goque

func Init() {
	gGoque = newGoque()
	go gGoque.Work()
}

func newGoque() *Goque {
	return &Goque{
		lq: NewQueue(),
		tk: time.NewTicker(time.Millisecond * 500),
	}
}

func GetGoque() *Goque {
	return gGoque
}

// Add 添加任务
func (g *Goque) Add(wf task.WorkerFunc, args, name string) {
	t := &task.Task{
		TaskId:      fmt.Sprintf("%d", time.Now().UnixNano()),
		ResultState: 2,
		WorkerName:  name,
		WorkerArgs:  args,
		WorkerFunc:  wf,
	}
	g.lq.Add(t)
}

func (g *Goque) check() {
	copyQueue := g.lq.CopyQueue()
	for taskInter := copyQueue.Top(); taskInter != nil; taskInter = copyQueue.Top() {
		task, ok := taskInter.(*task.Task)
		if !ok {
			log.Printf("[check] queue elem err, type is %s", reflect.TypeOf(task).Name())
			continue
		}
		task.WorkerStart = time.Now().Format("2006-01-02 15:04:05")

		// 起5个线程去上传
		if copyQueue.Length()%5 == 0 {
			g.wg.Wait()
		}

		g.wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer func() {
				task.WorkerEnd = time.Now().Format("2006-01-02 15:04:05")
				if err := model.WriteTaskToDB(task); err != nil {
					log.Println("[check] write db err: ", err)
				}
				g.wg.Done()
			}()
			res, err := task.WorkerFunc(task.WorkerArgs) // 执行任务
			if err != nil {
				task.ResultMsg = err.Error()
				return
			}
			task.ResultName = res
			task.ResultState = 1
			task.ResultUrl = "http://localhost:8080/media/video/" + res
		}(&g.wg)
	}
}

func (g *Goque) Work() {
	for {
		select {
		case <-g.tk.C:
			g.check()
		}
	}
}

// 打印当前队列
func (g *Goque) Dump() string {
	q, _ := json.Marshal(g.lq.buf)
	return string(q)
}
