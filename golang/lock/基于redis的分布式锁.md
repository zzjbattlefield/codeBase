### 基于redis的分布式锁
[redsync](https://github.com/go-redsync/redsync)

```go
package main

import (
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

func main() {
    //连接redis
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	rs := redsync.New(pool)

	//给锁命名
	mutexname := "my-global-mutex"
	mutex := rs.NewMutex(mutexname)

	if err := mutex.Lock(); err != nil {
		panic(err)
	}

	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}
}
```
## redis分布式锁要点
1. setnx可以保证获取和设置变成原子性操作;
2. 设置了锁但是还没有来得及解锁的时候服务挂掉了(死锁);
+ 1. 设置过期时间
+ 2. 设置了过期时间但是我业务还没执行完锁就过期了
    + 1. 在过期前手动刷新一下
    + 2. 需要自己启动协程完成刷新操作

## 分布式锁解决的问题
1. 互斥性 - setnx
2. 死锁 (超时机制)
3. 安全性
+ 1. 当时设置的值只有设置的g才能知道
+ 2. 在删除的时候取出redis中的值和自己保存的值对比一下