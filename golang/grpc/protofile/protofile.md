<font size="4"> 

# protobuf

## 生成protobuf文件的命令
```
protoc -I . protobuf文件.proto --go_out=plugins=grpc:.
```

## proto文件格式
```protobuf
//protobuf版本
syntax = "proto3";
//生成文件的路径eg: go_package = "/common/stream/proto/v1" 会生成到指定目录包名为上层的文件夹名v1
//生产到当前目录用. 包名用;分开
option go_package = ".;proto";
//引入底下的base.proto文件
import "base.proto";

service Hello {
    rpc SayHello(HelloRequest) returns (HelloRespone);
}
```

## proto文件的import
### base.proto(此文件也需要生成.pb.go文件)
```protobuf
syntax = "proto3";
option go_package = ".;proto";
message HelloRequest{
    string Name = 1;
}

message HelloRespone {
    string Message = 1;
}
```

## proto枚举
base.proto
```protobuf
syntax = "proto3";
option go_package = ".;proto";
message HelloRequest{
    Gender g =2;
    string Name = 1;
}

message HelloRespone {
    
    string Message = 1;
}

enum Gender{
    MALE = 0;
    FEMALE = 1;
}
```
client.go
```go
    rsp, err := client.SayHello(context.Background(), &proto.HelloRequest{G: proto.Gender_MALE})
```

## proto map
```protobuf
message HelloRequest{
    Gender g =2;
    string Name = 1;
    map<string,string> mp = 3;
}
```
client.go
```go
rsp, err := client.SayHello(context.Background(), &proto.HelloRequest{Mp: map[string]string{"msg": "hello"}})
```

## proto 时间戳类型
```protobuf
import "google/protobuf/timestamp.proto";
message HelloRequest{
    Gender g =2;
    string Name = 1;
    map<string,string> mp =3;
    google.protobuf.Timestamp addTime = 4;
}
```
client.go
```go
import "google.golang.org/protobuf/types/known/timestamppb"
rsp, err := client.SayHello(context.Background(), &proto.HelloRequest{ AddTime: timestamppb.New(time.Now())})
```