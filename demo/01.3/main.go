package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	//"runtime/debug"
	"sync"
	"syscall"
	"time"
)

func task(ctx context.Context)  {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(time.Second)
		panic("unexpected error")
	}
}

func main()  {
	ctx, cancel := context.WithCancel(context.Background())

	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		defer w.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("task exit with error: %v\n", err)
				//debug.PrintStack()
			}
		}()
		task(ctx)
		fmt.Println("task exit")
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
	cancel()
	w.Wait()
	fmt.Println("main exit")
}
