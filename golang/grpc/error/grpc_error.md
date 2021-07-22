<font size="4">

# grpc的错误处理
## service端返回错误
service.go
```go
    func (s *HelloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloRespone, error) {
	return nil, status.Error(codes.NotFound, "参数未找到")
    //或者
    return nil,status.Errorf(codes.NotFound,"参数未找到:%s",req.Name)
}
```
## client端接收错误
client.go
```go
    rsp, err := client.SayHello(ctx, &proto.HelloRequest{})
	if err != nil {
        //将error类型转成grpc.status类型获取错误具体数据
		s, ok := status.FromError(err)
		if !ok {
			panic("解析错误失败")
		}
		fmt.Printf("错误码:%d,错误信息:%s,", s.Code(), s.Message())
	} else {
		fmt.Println(rsp)
	}
```