package main

import "fmt"
// START1 OMIT
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums { out <- n }
		close(out)
	}()
	return out
}
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in { out <- n * n}
		close(out)
	}()
	return out
}
func main() {
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n) // 16 then 81
	}
}
// STOP1 OMIT