package main

import "fmt"

// START0 OMIT
func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; ; i ++ {
			c <- i
		}
	}()

	return c
}
// STOP0 OMIT

func main() {
	j := 1
// START1 OMIT
	c := gen()
	for i := range c {
		fmt.Println(i)
// STOP1 OMIT
		if j > 3 {
			return
		}
		j++		
	}
}