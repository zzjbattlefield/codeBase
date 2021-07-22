<font size="4"> 

# grpc拦截器
## service端
service.go
```go
    func main() {
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("你进入拦截器")
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			log.Println("no metaData")
		}
		for mdKey, mdData := range md {
			fmt.Printf("key is: %s,val is:%s", mdKey, mdData)
		}
        //执行后续的方法
		rsp, err := handler(ctx, req)
		fmt.Println("你退出拦截器")
		return rsp, err
	}
    //使用UnaryInterceptor生成一个grpc的option将option放入NewService中
    // UnaryInterceptor方法接受一个UnaryServerInterceptor
    //UnaryServerInterceptor的定义: func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
	opt := grpc.UnaryInterceptor(interceptor)
	s := grpc.NewServer(opt)
	proto.RegisterHelloServer(s, &HelloService{})
	listen, _ := net.Listen("tcp", "0.0.0.0:50091")
	err := s.Serve(listen)
	if err != nil {
		panic("faild to start grpc: " + err.Error())
	}
}
```

## client端
clinet.go
和service端是相似的逻辑
```go
interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		fmt.Println("客户端拦截器")
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Println(time.Since(start))
		return err
	}
	opt := grpc.WithUnaryInterceptor(interceptor)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure(), opt)
```