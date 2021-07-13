package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct{}

func main() {
	//可以携带参数的context
	c := context.WithValue(context.Background(), paramKey{}, "abc")
	//在携带参数的context的基础上添加timeOut 手动cancel或五秒后所有基于c的context都会退出
	c, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()
	mainTask(c)
	fmt.Println("done")
	time.Sleep(time.Hour)
}

func mainTask(ctx context.Context) {
	fmt.Printf("mainTask start with paramKey:%q\n", ctx.Value(paramKey{}))
	smallTask(ctx, "task1")
	smallTask(ctx, "task2")
}

func smallTask(ctx context.Context, name string) {
	fmt.Printf("%s started with paramKey:%q\n", name, ctx.Value(paramKey{}))
	select {
	case <-time.After(6 * time.Second):
		fmt.Printf("%s finsh \n", name)
	case <-ctx.Done():
		fmt.Printf("%s cancelled \n", name)
	}
}
