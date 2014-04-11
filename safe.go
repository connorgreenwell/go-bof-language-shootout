package main

import (
	"fmt"
	"time"
)

// START SAFE OMIT
func GoPrint(pre string, ch <-chan int) {
	for {
		fmt.Println(pre, <-ch)
	}
}

func main() {
	ch := make(chan int)
	go GoPrint("A", ch)
	go GoPrint("B", ch)
	go GoPrint("C", ch)
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(250 * time.Millisecond)
	}
}

// END SAFE OMIT
