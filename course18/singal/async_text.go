package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go doYourTask()

	sig := <-sigChan
	fmt.Printf("signal: %v\n", sig)

	cleanup()

	os.Exit(0)
}

func doYourTask() {
	println("111")
	time.Sleep(3 * time.Second)
	println("333")
}

func cleanup() {
	time.Sleep(3 * time.Second)
	println("222")
}
