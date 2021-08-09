<font size="4">

# workerPool
## work.go
```go
    package main

import (
	"fmt"
	"sync"
)

type Pool struct {
	Name    string //池名称
	Workers []*Worker 
	Size    int //池内worker数量

	JobQueue  chan Job //工作队列
	QueueSize int //队列长度
}
//wokerPool初始化方法 返回初始化pool的指针
func NewPool(name string, size int, QueueSize int) *Pool {
	pool := &Pool{
		Name:      name,
		Size:      size,
		QueueSize: QueueSize,
	}
	if size < 1 {
		pool.Size = 1
	}
	if QueueSize < 1 {
		pool.QueueSize = 1
	}
	pool.Workers = make([]*Worker, size)
	for i := 0; i < pool.Size; i++ {
		worker := Worker{Name: fmt.Sprintf("%s-worker-%d", pool.Name, i+1), ID: i + 1}
		pool.Workers[i] = &worker
	}
	pool.JobQueue = make(chan Job, pool.QueueSize)
	return pool
}

//循环调用每个池子里的worker的Start()方法
func (p *Pool) Start() {
	for _, worker := range p.Workers {
		worker.Start(p.JobQueue)
	}
	fmt.Println("all worker is start")
}

//Stop方法直接关闭了Job队列,worker检查队列关闭会直接退出 worker.Stop仅打印一下
func (p *Pool) Stop() {
	close(p.JobQueue)
	var wg sync.WaitGroup
	for _, worker := range p.Workers {
		wg.Add(1)
		go func(w *Worker) {
			defer wg.Done()
			w.Stop()
		}(worker)
	}
	wg.Wait()
	fmt.Println("all worker stop")
}

```

## woker.go
```go
package main

import (
	"fmt"
)

type Worker struct {
	StopChan chan bool //用来判断worker是否关闭
	Name     string
	ID       int
}

func (w *Worker) Start(jobQueue chan Job) {
	w.StopChan = make(chan bool)
	go func() {
		for {
			if job, ok := <-jobQueue; ok {
				job.Start(w) // worker执行任务
			} else {
				//jobChannel Close
				fmt.Printf("worker:%s stop", w.Name)
				w.StopChan <- true
				break
			}
		}

	}()
}

func (w *Worker) Stop() {
	<-w.StopChan
	fmt.Printf("worker %s stopped\n", w.Name)
}

```

## main.go
```go
package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
    //确保退出前所有goroutine都已经退出
	wg := sync.WaitGroup{}
	wg.Add(1)
	defer func() {
		wg.Wait()
		fmt.Println("numGorutine:", runtime.NumGoroutine())
	}()
	pool := NewPool("testPool", 5, 20)
	pool.Start()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	stopChan := make(chan struct{})
	defer cancel()
	go func() {
        //单独一个协程监听context事件
        //当context取消了的时候给stopChan信号 不再往队列里加入新的任务 跳出循环
		<-ctx.Done()
		stopChan <- struct{}{}
		log.Println(ctx.Err().Error())
		pool.Stop()
		wg.Done()
	}()
Loop:
	for i := 0; i < 500; i++ {
		select {
		case <-stopChan:
			break Loop
		default:
			job := PrintJob{Index: i + 1}
			pool.JobQueue <- job
		}
	}
}

type PrintJob struct {
	Index int
}

type Job interface {
	Start(*Worker) (interface{}, error)
}

func (j PrintJob) Start(w *Worker) (interface{}, error) {
	fmt.Printf("job %s - %d \n", w.Name, j.Index)
	return nil, nil
}

```