package main

import (
	"fmt"
	"sync"
	"time"
)

func task1()  {
	time.Sleep(time.Second)
}

func task2()  {
	time.Sleep(time.Second * 2)
}

func main()  {
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		defer w.Done()
		task1()
		fmt.Println("task1 exit")
	}()
	w.Add(1)
	go func() {
		defer w.Done()
		task2()
		fmt.Println("task2 exit")
	}()
	w.Wait()
	fmt.Println("main exit")
}
