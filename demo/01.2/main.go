package _1_2

import (
	"context"
	"fmt"
	"os"
	"os/signal"
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
	}
}

func main()  {
	ctx, cancel := context.WithCancel(context.Background())

	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		defer w.Done()
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
