package main

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
)

func main() {
	url := fmt.Sprintf("%s", "127.0.0.1:4150")
	conf := nsq.NewConfig()
	producer, err := nsq.NewProducer(url, conf)
	ch := make(chan struct{})
	for i := 0; i < 100000; i++ {
		go func() {
			if err != nil {
				println(err.Error())
				return
			}
			seed := &LocalTest{
				Url: "https://pan.baidu.com/share/init?surl=tewYeydWMU85D-KM_Ti4lg&pwd=uanw",
			}
			seedBody, _ := json.Marshal(seed)
			producer.Publish("local.Test.Url", seedBody)
			println("send")
		}()
	}
	<-ch
}
