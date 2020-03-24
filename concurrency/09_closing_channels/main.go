package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	//Even if closing can be possible in the receiver, it will error out if 
	// close() is defined in a reciever function
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		fmt.Println(<-ch) //will print 0 if you send only 1 item in this example
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		close(ch) //Cannot open a channel after it's closed
					//Closing needs to be always in the sending side of the operation
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
