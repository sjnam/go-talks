package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	var c <-chan int
	select {
	case <-c: // HL
	case <-time.After(1*time.Second):
		fmt.Println("Timed out.")
	}
	// STOP OMIT
}
