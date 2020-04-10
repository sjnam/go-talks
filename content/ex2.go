// +build ignore,OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(2 * time.Second)
		close(c) // HL
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c: // HL
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
	// STOP OMIT
}
