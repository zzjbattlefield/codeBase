package main

import (
	"fmt"
	"time"
)

//协程池

type Pool struct {
	work chan func()
	sem  chan struct{}
}

func NewPool(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.sem <- struct{}{}:
		go p.worker(task)
	}
}

func (p *Pool) worker(task func()) {
	defer func() {
		<-p.sem
	}()
	for {
		task()
		task = <-p.work
	}
}

func main() {
	pool := NewPool(32)
	for i := 0; i < 5; i++ {
		count := i
		pool.NewTask(
			func() {
				fmt.Println("start task", count)
				time.Sleep(time.Second)
				fmt.Println("finsh task", count)
			},
		)
	}
	time.Sleep(5 * time.Second)
}
