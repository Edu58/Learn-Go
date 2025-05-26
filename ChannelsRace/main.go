package main

import "fmt"

func printer(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}

	close(ch)
}

func main() {
	ch := make(chan bool)
	go printer(ch, 5)

	for val := range ch {
		fmt.Printf("%v\t", val)
	}

	for i := 0; i < 15; i++ {
		fmt.Printf("%v\t", <-ch)
	}
}
