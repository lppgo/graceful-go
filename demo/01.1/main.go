package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func task1(ctx context.Context)  {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		time.Sleep(time.Second)
	}
}

func task2(ctx context.Context)  {
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
		task1(ctx)
		fmt.Println("task1 exit")
	}()
	w.Add(1)
	go func() {
		defer w.Done()
		task2(ctx)
		fmt.Println("task2 exit")
	}()
	w.Add(1)
	go func() {
		defer w.Done()
		time.Sleep(time.Second * 3)
		fmt.Println("cancel")
		cancel()
	}()
	w.Wait()
	fmt.Println("main exit")
}
