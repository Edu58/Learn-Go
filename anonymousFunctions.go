package main

import "fmt"

func main() {
	func(i int) (s int) {
		fmt.Println(i * i)
		return
	}(4)
}
