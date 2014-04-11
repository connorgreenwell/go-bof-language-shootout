package main

import (
	"fmt"
	"time"
)

func WaitThenPrint(msg string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(msg)
}

// START GOSLOW OMIT
func main() {
	go WaitThenPrint("hello")
	fmt.Println("world")
	time.Sleep(1 * time.Second)
}

// END GOSLOW OMIT
