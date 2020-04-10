package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	done := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
		}
		workCounter++
		time.Sleep(time.Second)
	}
	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
	// STOP OMIT
}


	