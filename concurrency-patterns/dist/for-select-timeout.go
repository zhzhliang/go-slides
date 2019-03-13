package main

import (
	"fmt"
	"time"
)

func loop(ch chan <- int) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go loop(ch)
	// START0 OMIT
	timeout := time.After(3 * time.Second)
	for {
		select {
		case s := <- ch:
			fmt.Println(s)
		case <- timeout:
			fmt.Println("timeout")
			return
		}
	}
	// STOP0 OMIT
}