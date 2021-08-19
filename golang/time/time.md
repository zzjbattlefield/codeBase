<font size="4">

# 标准库time包

## 时间类型
time.Time表示时间类型 通过time.now获取当前时间对象 可以获取对象的年月日时分秒信息
```go
func main() {
	now := time.Now() //time.Time
	now.Year()        //年
	now.Month()       //月
	now.Day()         //日
	now.Hour()        //小时
	now.Minute()      //分钟
	now.Second()      //秒
}
```

## 时间戳
```go
func main(){
    now := time.Now()       //time.Time
	timeStamp := now.Unix() //时间戳
	now.Nanosecond()        //纳秒时间戳

	time.Unix(timeStamp, 0) //将时间戳转成时间格式 //time.Time

}
```
## 时间间隔
```go
time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。
time.Duration表示一段时间间隔，可表示的最长时间段大约290年
time包中定义的时间间隔类型的常量如下：
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)

例如：time.Duration表示1纳秒，time.Second表示1秒。
```

### 时间操作

#### 时间操作Add
```go
   //func (t Time) Add(d Duration) Time 
   func main(){
       now := time.Now() //time.Time
	    later := now.Add(60 * time.Minute)
   }
```

#### 时间操作Sub

```go
    //func (t Time) Sub(u Time) Duration
	func main(){
        now := time.Now() //time.Time
        later := now.Add(60 * time.Minute)
        fmt.Println(later.Sub(now))
    }

```

## 定时器
使用time.Tick(时间间隔)来设置定时器，定时器的本质上十一个通道(channel)。
```go
    func main(){
        ticker := time.Tick(time.Second)
        i := 1
        for range ticker {
            fmt.Printf("第%d秒", i)
            i++
        }

        time.
    }
```
如果在for循环里调用time.Tick会造成协程泄露
```go
// func After(d Duration) <-chan Time 接受一个时间段 到期后发送值到channel
func main() {
	after := time.After(time.Second)
	<-after
	fmt.Println("一秒钟到了")
    //func AfterFunc(d Duration, f func()) *Timer 接受一个时间段和一个匿名函数 到期后调用匿名函数
    time.AfterFunc(time.Second, func() { fmt.Println("一秒钟到了") })
	time.Sleep(time.Second * 2)
}
```

## 时间的格式化
时间类型有一个自带的方法Format进行格式化，格式化时间模板是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）
```go
//func (t Time) Format(layout string) string 
now := time.Now()
// 24小时制
fmt.Println(now.Format("2006-01-02 15:04:05"))//2021-08-19 17:01:33
// 12小时制
fmt.Println(now.Format("2006-01-02 03:04:05 PM"))//2021-08-19 05:01:33 PM
fmt.Println(now.Format("2006/01/02 15:04"))
fmt.Println(now.Format("15:04 2006/01/02"))
fmt.Println(now.Format("2006/01/02"))
```

## 解析字符串格式的时间

```go
// ParseInLocation(layout, value string, loc *Location) (Time, error)
time, err := time.ParseInLocation("2006/01/02 15:04:05", "2021/01/02 13:14:15", time.Local)
if err != nil {
    panic(err)
}
fmt.Println(time)

```