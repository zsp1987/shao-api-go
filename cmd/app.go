package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
	}()

	go func() {
		chan2 <- 2
	}()

	select {
	case <-chan1:
		fmt.Println("chan1")
	case <-chan2:
		fmt.Println("chan2")
	default:
		fmt.Println("default")
	}
}
