package main

import (
	"fmt"
	"sync"
)

func work(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		//go func() { done <- true }()
		wg.Done()
	}
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("worker %d received %c\n", id, n)
	//}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go work(id, w.in, wg)
	return w
}

func chanDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	wg.Wait()
}

func main() {
	//channelClose()
	//bufferedChannel()
	chanDemo()
}
