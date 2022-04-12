# GO server&client for GRPC

## 生成PHP源码

1. 安装 
     + protoc 根据操作系统具体安装：https://github.com/protocolbuffers/protobuf/tags
     + protoc-gen-go | go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
     + protoc-gen-go-rpc | go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

2. pb协议(demo.proto)内容

``` pb
	
	//版本
	syntax = "proto3";

	//服务名
	package Hello;

	//包名
	option go_package = "example/grpc";

	//定义服务名及方法
	service Demo {
	    rpc GetDemo (GetDemoReq) returns (GetDemoReply) {}
	}

	//定义请求消息
	message GetDemoReq {
	    int64 user_id = 1;
	}

	//定义响应消息
	message GetDemoReply {
	    int64 user_id = 1;
	}

```

3. 执行  

	+ protos -I=协议目录 --php_out=生成目录 --go-grpc_out=grpc生成目录  协议文件

	+ protoc -I=./protos --go_out=. --go-grpc_out=. demo.proto
	
	

4. 生成目录

``` go   
   
│  go.mod
│  go.sum
│  README.md
│
├─example
│  └─grpc
│        demo.pb.go
│        demo_grpc.pb.go
│
└─protos
        demo.proto

```
    

## 包依赖
``` json
 go get google.golang.org/grpc
 go mod tidy
 
```
## RPC调用


``` go

package main

import (
	"context"
	pb "gogrpc/example"
	"google.golang.org/grpc"
	"log"
	"net"
	"flag"
)

/*const (
	port = ":50051"
)*/


var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedDemoServer
}

func NewServer() pb.DemoServer {
	return &server{}
}

func (s *server) GetDemo(ctx context.Context, in *pb.GetDemoReq) (*pb.GetDemoReply, error) {
	log.Printf("Received: %d", in.UserId)
	return &pb.GetDemoReply{UserId: in.UserId}, nil
}

func main() {

	flag.Parse()
	// lis, err := net.Listen("tcp", port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//pb.RegisterDemoServer(s, &server{})
	pb.RegisterDemoServer(s, NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

	
	

```