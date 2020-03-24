package main

import (
	"fmt"
	"time"
	"sync"
	"runtime" //Only to enable parallelism
)


func main() {
	
	runtime.GOMAXPROCS(2) //Only to enable parallelism
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	//Anonymous function
	//jhaving a parenthesis at the end of the function makes it self executing
	go func(){
		defer waitGrp.Done()

		time.Sleep(5*time.Second)
		fmt.Println("Hello")
	}()

	go func(){
		defer waitGrp.Done()

		fmt.Println("DX")
	}()

	//Wait for the go routines to finish.
	// without the waitGrp statement, main finishes before goroutines are done
	waitGrp.Wait()



}
