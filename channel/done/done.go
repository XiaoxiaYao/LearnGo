package main

import (
	"fmt"
	"sync"
)

func work(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		//go func() { done <- true }()
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go work(id, w)
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
