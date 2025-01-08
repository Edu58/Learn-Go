package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new ticker with a 10-second period
	ticker := time.NewTicker(3 * time.Second)

	// Create a channel to signal when to stop the ticker
	quit := make(chan struct{})

	// Run a goroutine that listens to the ticker and the quit channel
	go func() {
		for {
			select {
			case <-ticker.C: // On each tick, run the task
				fmt.Println("Running task...")
				// Do your task here
			case <-quit: // On quit signal, stop the ticker and return
				ticker.Stop()
				return
			}
		}
	}()

	// Simulate some work for 30 seconds
	time.Sleep(100 * time.Second)

	// Send a quit signal to stop the ticker
	// close(quit)

	// Print a message when the ticker is stopped
	fmt.Println("Ticker is stopped!")
}
