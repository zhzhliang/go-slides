// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START1 OMIT
func main() {
	cs := fanOut(boring("test"), 3)
	for i := 0; i < 3; i++ {
		s := <- cs[i]
		fmt.Println(s)
	}
}
// STOP1 OMIT

// START2 OMIT
func boring(msg string) <-chan string { // Returns receive-only channel of strings. // HL
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function. // HL
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller. // HL
}
// STOP2 OMIT


// START3 OMIT
func fanOut(in <-chan string, n int) []<-chan string {
	cs := make([]chan string, 0, n)
	for i := 0; i < n; i ++ {
		cs = append(cs, make(chan string))
	}
	go func () {
		// defer close channels omitted
		for _, c := range cs {
			select {
			case v, ok := <- in:
				if !ok { return }
				c <- v
			}
		}
	}()
	rcs := make([]<-chan string, 0, n)
	for i := 0; i < n; i++ {
		rcs = append(rcs, cs[i])
	}
	return rcs
}
// STOP3 OMIT

