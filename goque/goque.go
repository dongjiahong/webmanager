package goque

import (
	"encoding/json"
	"log"
	"reflect"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Goque struct {
	lq *Queue // 任务接受队列 loop queue
	dq *Queue // 任务处理队列 done queue // XXX 这些结果应该本应该放在redis里等地方
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
		dq: NewQueue(),
		tk: time.NewTicker(time.Millisecond * 500),
	}
}

func GetGoque() *Goque {
	return gGoque
}

// Add 添加任务
func (g *Goque) Add(wf workerFunc, args string) {
	id := uuid.Must(uuid.NewV4()).String()
	t := &Task{
		Id:          id,
		ResultState: 2,
		Worker: &Worker{
			wf:   wf,
			args: args,
		},
	}
	g.lq.Add(t)
}

func (g *Goque) check() {
	copyQueue := g.lq.CopyQueue()
	for taskInter := copyQueue.Top(); taskInter != nil; taskInter = copyQueue.Top() {
		task, ok := taskInter.(*Task)
		if !ok {
			log.Printf("[check] queue elem err, type is %s", reflect.TypeOf(task).Name())
			continue
		}

		// 起5个线程去上传
		if copyQueue.Length()%5 == 0 {
			g.wg.Wait()
		}

		g.wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer g.wg.Done()
			w := task.Worker
			res, err := w.wf(w.args) // 执行任务
			if err != nil {
				res = err.Error()
			}
			task.Result = res
			g.dq.Add(task)
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

// 打印结果队列
func (g *Goque) DumpDone() string {
	q, _ := json.Marshal(g.dq.buf)
	return string(q)
}
