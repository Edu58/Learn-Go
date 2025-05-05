package runes

import (
	"fmt"
	"time"
)

func test_rune() {
	var randString = []rune("489234")
	fmt.Println(string(randString))
	fmt.Println(time.Now().UnixNano())
}
