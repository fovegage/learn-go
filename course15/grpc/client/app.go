package main

import (
	"context"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
	pb "study/course15/grpc/proto"
	"time"
)

func main() {
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatal(err)
	}

	// 配置一条限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test",
			Threshold:              10000,
			TokenCalculateStrategy: flow.WarmUp,
			ControlBehavior:        flow.Throttling,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := grpc.Dial("127.0.0.1:8789", grpc.WithTransportCredentials(insecure.NewCredentials()))

	ch := make(chan struct{})
	for i := 0; i < 1000; i++ {
		go func() {
			e, b := sentinel.Entry("some-test")
			if b != nil {
				println("error")
			} else {
				// 请求允许通过，此处编写业务逻辑
				fmt.Println(util.CurrentTimeMillis(), "Passed")
				time.Sleep(time.Duration(rand.Uint64()%10) * time.Millisecond)
				DoRequest(conn)
				e.Exit()
			}
		}()
	}
	<-ch
}

func DoRequest(conn *grpc.ClientConn) {
	//conn, err := grpc.Dial("127.0.0.1:8789", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	return
	//}
	//defer conn.Close()
	ctx := context.Background()
	hello := pb.NewGreeterClient(conn)
	req := &pb.HelloRequest{
		Name: "Gage",
	}
	stream, err := hello.SayHello(ctx, req)
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			println("eof")
			break
		}
		if err != nil {
			log.Printf("failed to recv: %v", err)
		}
		log.Printf("Greeting: %s", reply.Message)
	}
	println("111")
}
