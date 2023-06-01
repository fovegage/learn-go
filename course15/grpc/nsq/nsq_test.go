package main

import (
	"context"
	"encoding/json"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/nsqio/go-nsq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"study/course15/grpc/lock"
	pb "study/course15/grpc/proto"
	"testing"
)

// 模拟nsq

func TestNewApp(t *testing.T) {
	var (
		conf = nsq.NewConfig()
	)
	// 加载限流规则
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatal(err)
	}

	// 配置一条限流规则 并发1000测试
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "crawl-test",
			Threshold:              10,
			TokenCalculateStrategy: flow.WarmUp,
			ControlBehavior:        flow.Throttling,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan struct{})
	//conf.MaxAttempts = nsq2.MaxAttempts
	//conf.MaxInFlight = 1

	loopupdAddr := "127.0.0.1:4161"

	// url
	productUrl, err := nsq.NewConsumer("local.Test.Url", "pipeline", conf)

	conn, err := grpc.Dial("127.0.0.1:8789", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// 设置消息处理函数
	productUrl.AddConcurrentHandlers(&ConcurrentTestHandlers{itemLock: lock.NewKeyLock(), conn: conn}, conf.MaxInFlight)
	// nsqlookupd
	if err = productUrl.ConnectToNSQLookupds([]string{loopupdAddr}); err != nil {
		log.Fatal(err)
	}
	<-ch
	go func() {
		productUrl.Stop()
	}()
}

type ConcurrentTestHandlers struct {
	itemLock *lock.KeyLock
	conn     *grpc.ClientConn
}

func (c *ConcurrentTestHandlers) HandleMessage(message *nsq.Message) error {
	message.DisableAutoResponse()
	var (
		newUrl LocalTest
		err    error
	)
	err = json.Unmarshal(message.Body, &newUrl)
	if err != nil {
		message.Finish()
		return err
	}
	println(newUrl.Url)

	// todo: url状态数据插入数据库
	// 基于协程的统计
	//e, b := sentinel.Entry("crawl-test")
	//if b != nil {
	//	log2.BuildLog("sentinel", fmt.Sprintf("send mq url %s", b.Error()), "")
	//} else {
	//	fmt.Printf("crawl-api recv url %s\n", newUrl.Url)
	//	DoRequest(c.conn)
	//	if err != nil {
	//		message.Finish()
	//		log2.BuildLog("api", fmt.Sprintf("send mq url %s", err.Error()), "")
	//	}
	//	e.Exit()
	//}
	message.Finish()
	return nil
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
