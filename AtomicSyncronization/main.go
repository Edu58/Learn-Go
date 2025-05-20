package main

import (
	"fmt"
	"sync/atomic"
)

// The WaitGroup uses a single atomic.Uint64 (wg.state) to atomically store and update two values:
// 1.	High 32 bits: the counter — number of goroutines that are running (Add/Done logic).
// 2.	Low 32 bits: the waiter count — number of goroutines blocked in Wait.
type Counter struct {
	state atomic.Uint64
}

func main() {
	var c Counter

	c.IncrementActive()
	c.IncrementActive()
	c.DecrementFailed()

	active, failed := c.Snapshot()
	fmt.Printf("Active: %v\n", active)
	fmt.Printf("Failed: %v\n", failed)
}

func (c *Counter) IncrementActive() {
	c.state.Add(1 << 32)
}

func (c *Counter) DecrementFailed() {
	c.state.Add(1)
}

func (c *Counter) Snapshot() (active uint32, failed uint32) {
	val := c.state.Load()
	active = uint32(val >> 32)
	failed = uint32(val)
	return
}
