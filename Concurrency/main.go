package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// An unbuffered channel will block when sending if there is no receiver
	c := make(chan int, 1)

	wg.Add(1)
	go counter(c)
	wg.Wait()
	output := <-c
	fmt.Println("OUTPUT IS: ", output)
}
