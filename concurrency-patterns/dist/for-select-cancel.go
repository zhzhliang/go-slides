package main

import (
	"fmt"
	"time"
)
// START0 OMIT
func loop(cancel <-chan interface{}) {
	for i := 0; ; i++{
		select {
		case <- cancel:
			return
		default:
		}
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
// STOP0 OMIT

func main() {
	cancel := make(chan interface{})

	go loop(cancel)
	time.Sleep(3 * time.Second)
	cancel <- 1
	
	time.Sleep(5 * time.Second)
}