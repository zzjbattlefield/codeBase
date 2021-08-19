<font size="4"> 

# GRPC客户端
## 连接服务端
```go
package main

func main() {
	var host string
	var port int
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewHelloClient(conn)
	rsp, err := client.SayHello(context.Background(), &proto.HelloRequest{Name: "hello grpc"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rsp)
}
```