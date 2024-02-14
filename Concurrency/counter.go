package main

import (
	"fmt"
)

func counter(c chan int) {
	defer wg.Done()
	sum := 0

	for i := 0; i < 100; i++ {
		fmt.Println("IDX From First Func: ", i)
		sum += i
	}

	c <- sum
}
