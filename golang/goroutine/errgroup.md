<font size="4">

## 什么是errgroup
errgroup 在 WaitGroup 的基础上实现子协程错误传递, 同时使用 context 控制协程的生命周期。

## errgroup的结构
```go
type Group struct {
    cancel  func()             //context cancel()
    wg      sync.WaitGroup         
    errOnce sync.Once          //只会传递第一个出现错的协程的 error
    err     error              //传递子协程错误
}
//创建一个带有cancel方法的Group结构体
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}
//将要用go并发的方法使用此函数调用
func (g *Group) Go(f func() error) {
    g.wg.Add(1)

    go func() {
        defer g.wg.Done()
        if err := f(); err != nil {
            g.errOnce.Do(func() {       
                g.err = err             //记录子协程中的错误
                if g.cancel != nil {
                    g.cancel()
                }
            })
        }
    }()
}
//调用此方法等待所有goroutine方法完成 返回err 如果context中有cancel方法会调用
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}

```

```go
type query struct{
	url string
	err error
}
func main() {
	resultChan := make(chan query)
	urls := []query{
		{url: "http://www.baidu.com"},
		{url: "http://www.google.com"},
		{url: "http://www.sina.com"},
	}
	go func() {
		for r := range resultChan {
			if r.err != nil {
				log.Printf("%s error: %v", r.url, r.err)
				continue
			}
			log.Println(r.url)
		}
	}()
	err := fetch(urls, resultChan)
	if err != nil {
		log.Println(err.Error())
	}
}

func fetch(urls []query, resultChan chan query) error {
	eg, ctx := errgroup.WithContext(context.Background())
	for _, url := range urls {
		url := url
		eg.Go(func(ctx context.Context, url query, resultChan chan query) func() error {

			return func() error {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				defer cancel()
				req, err := http.NewRequestWithContext(ctx, "GET", url.url, nil)
				if err != nil {
					url.err = err
					resultChan <- url
					return err
				}
				_, err = http.DefaultClient.Do(req)
				if err != nil {
					url.err = err
					resultChan <- url
					return err
				}
				resultChan <- url
				return nil
			}

		}(ctx, url, resultChan))
	}
	return eg.Wait()
}
```
