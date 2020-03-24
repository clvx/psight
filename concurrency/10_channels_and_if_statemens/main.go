package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)	
		} else {
			fmt.Println("channel not open")
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0 //prints 0, true
		//close(ch) /prints channel not open, if previous is commented out
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
