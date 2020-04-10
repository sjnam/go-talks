// +build ignore,OMIT

package main

func main() {
	// START OMIT
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}
	select { // HL
	case <-c1:
		// Do something
	case <-c2:
		// Do something
	case c3 <- struct{}{}:
		// Do something
	} // HL
	// STOP OMIT
}
