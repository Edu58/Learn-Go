package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)

		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("Running task %d \n", x)
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("I'm Done")
}
