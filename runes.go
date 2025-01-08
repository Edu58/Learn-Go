package main

import (
	"fmt"
	"time"
)

func main() {
	var randString = []rune("489234")
	fmt.Println(string(randString))
	fmt.Println(time.Now().UnixNano())
}
