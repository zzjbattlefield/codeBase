<font size="4">

## 1.下面的代码会报错吗 为什么

```go
    type Param map[string]interface{}

    type Show struct {
        Param
    }

    func main() {
        s := new(Show)
        s.Param["RMB"] = 10000
    }
```
**解析**
 + new关键字无法初始化结构体中的属性 应该改成
 ```go
     func main() {
        s := new(Show)
        s.Param = make(Param)
        s.Param["RMB"] = 10000
    }
 ```

## 2.以下代码有什么问题
```go 
type People struct {
	name string `json:"name"`
}

func main() {
	js := `{
		"name":"11"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}
```
**解析**
 + 结构体内的属性是私有属性 json包无法进行私有属性的转换

## 3.以下代码有什么问题
```go
type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
 	p := &People{}
	p.String()
}
```
**解析**
 + p的String方法实际实现了String的接口 在fmt包中如何类型实现了string接口则会直接调用 上面会形成循环引用

## 4.找出以下代码的问题
```go
func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
```
**解析**
 + close方法可能会在起完两个协程后直接被调用 协程中往close的channel写数据会panic
 + 修改:在生产者close 使用waitGroup

```go 
func main() {
	ch := make(chan int, 1000)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				wg.Done()
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	wg.Wait()
	fmt.Println("ok")
}
```

## 5.下面代码，执行时为什么会报错
```go
type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"zhoujielun"}}
	m["people"].name = "wuyanzu"
}
```
**解析**
 + map的val本身是不可寻址所以不能直接赋值 可以使用临时变量或者直接存地址来解决
```go
type Student struct {
name string
}

func main() {
	m := map[string]Student{"people": {"zhoujielun"}}
	tmp := m["people"]
	tmp.name = "wuyanzu"
	//或者
	m1 := make(map[string]*Student)
	m1["name"] = &Student{"zhoujielun"}
	m1["name"].name = "wuyanzu"
}
```
