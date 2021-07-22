<font size="4"> 

# GRPC服务端
## 监听grpc服务
```go
package main

import (
	"context"
	"golangTest/grpc_test/proto"
	"net"

	"google.golang.org/grpc"
)

type HelloService struct{}

func main() {
    //先New grpcService
	s := grpc.NewServer()
    //注册结构体
	proto.RegisterHelloServer(s, &HelloService{})
    //监听客户端请求
	listen, _ := net.Listen("tcp", "0.0.0.0:50091")
    //提供服务
	err := s.Serve(listen)

	if err != nil {
		panic("faild to start grpc: " + err.Error())
	}
}

func (s *HelloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloRespone, error) {
	return &proto.HelloRespone{Message: "hello" + req.Name}, nil
}

```