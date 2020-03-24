package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	//As each goroutine is independent, the channel is the only way to communicate
	//between goroutines in this case.
	ch := make(chan int)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch) //receiving a message from channel
		wg.Done()
	}(ch, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42 //sending a message to channel
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
