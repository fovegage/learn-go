package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "study/course15/grpc/proto"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func (g *GreeterService) Parse(callback func(interface{})) {
	for i := 0; i < 10; i++ {
		callback(fmt.Sprintf("%d", i))
	}
}

func (g *GreeterService) SayHello(request *pb.HelloRequest, server pb.Greeter_SayHelloServer) error {
	//name := request.Name
	//for i := 0; i < 100; i++ {
	//	server.Send(&pb.HelloReply{Message: "Hello " + name + strconv.Itoa(i)})
	//}
	g.Parse(func(i interface{}) {
		println(i.(string))
		server.Send(&pb.HelloReply{Message: "Hello " + i.(string)})
	})
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8789")
	if err != nil {
		fmt.Printf("监听端口失败: %s", err)
		return
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &GreeterService{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("开启服务失败: %s", err)
		return
	}
}
