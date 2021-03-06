package main

import (
	"math/rand"
	"time"
	"fmt"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i:=0; i< 10; i++ {
		id := rnd.Intn(10) + 1
		if b, ok := queryCache(id); ok{
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			fmt.Println("from database")
			fmt.Println(b)
			continue
		}
		fmt.Printf(" Book not found with id: '%v'", id)
		time.Sleep(150 * time.Millisecond)
	}
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books{
		if b.ID == id{
			cache[id] = b
			return b, true
		}
	}

	return Book{}, false
}
