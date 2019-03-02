package goque

import (
	"sync"
)

const minQueueLen = 16

type Queue struct {
	buf               []interface{}
	l                 sync.RWMutex
	head, tail, count int
}

func NewQueue() *Queue {
	return &Queue{
		buf: make([]interface{}, minQueueLen),
	}
}

func (q *Queue) Length() int {
	q.l.RLock()
	defer q.l.RUnlock()

	return q.count
}

// 保内函数，在调用的函数里要注意上锁
func (q *Queue) resize() {

	newBuf := make([]interface{}, q.count<<1)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		//	        head
		//			 ^
		//	+-----+--+--------+
		//	|8|9| | |3|4|5|6|7|
		//	+--+--+-----------+
		//	   v
		//	   tail

		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

func (q *Queue) Add(elem interface{}) {
	q.l.Lock()
	defer q.l.Unlock()

	if q.count == len(q.buf) {
		q.resize()
	}

	q.buf[q.tail] = elem
	// 循环使用空间
	q.tail = (q.tail + 1) & (len(q.buf) - 1)
	q.count++
}

func (q *Queue) Peek() interface{} {
	q.l.RLock()
	defer q.l.RUnlock()

	if q.count <= 0 {
		return nil
	}
	return q.buf[q.head]
}

func (q *Queue) Get(i int) interface{} {
	q.l.RLock()
	defer q.l.RUnlock()

	if i < 0 {
		i += q.count
	}
	if i < 0 || i >= q.count {
		return nil
	}
	return q.buf[(q.head+i)&(len(q.buf)-1)]
}

func (q *Queue) Top() interface{} {
	t := q.Peek()
	if t != nil {
		q.Remove()
	}
	return t
}

func (q *Queue) CopyQueue() *Queue {
	copy := NewQueue()
	for i := 0; i < q.Length(); i++ {
		if elem := q.Top(); elem != nil {
			copy.Add(elem)
		}
	}
	return copy
}

func (q *Queue) Remove() interface{} {
	q.l.Lock()
	defer q.l.Unlock()

	if q.count <= 0 {
		return nil
	}
	ret := q.buf[q.head]
	q.buf[q.head] = nil

	q.head = (q.head + 1) & (len(q.buf) - 1)
	q.count--

	if len(q.buf) > minQueueLen && (q.count<<2) == len(q.buf) {
		q.resize()
	}
	return ret
}
