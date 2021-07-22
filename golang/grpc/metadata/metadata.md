<font size="4"> 

# MetaData
## 什么是MetaData
metadata是以key-value的形式存储数据的，其中key是string类型，而value是[]string，即一个字符串数组类型。metadata使得client和server能够为对方提供关于本次调用的一些信息，就像一次http请求的RequestHeader和ResponseHeader一样。http中header的生命周周期是一次http请求，那么metadata的生命周期就是一次RPC调用。

## 如何新建metadata
```go
import "google.golang.org/grpc/metadata"
//1使用metadata.New()
md := metadata.New(map[string]string{
		"name_md":     "grpc",
		"password_md": "123",
	})
//2.使用metadata.Pairs key和val用,隔开
md := metadata.Pairs(
    "key1", "val1",
    "key1", "val1-2", // "key1" will have map value []string{"val1", "val1-2"}
    "key2", "val2",
)

```

## 发送metadata
```go
ctx := metadata.NewOutgoingContext(context.Background(), md)
rsp, err := client.SayHello(ctx, &proto.HelloRequest{})
```

## 接受metadata
```go
md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metaData")
	}
	for mdKey, mdData := range md {
		fmt.Printf("key is: %s,val is:%s", mdKey, mdData)
	}
```