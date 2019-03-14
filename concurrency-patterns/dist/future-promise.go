package main

import "fmt"

// START1 OMIT
type Future chan string

func (ft Future)Get() string {
	return <-ft
}

func (ft Future)Listen(f func(string)) {
	go func() {
		v := <- ft
		f(v)
	}()
}
// STOP1 OMIT

// START2 OMIT
type Promise struct {
	ft Future
}
func NewPromise() *Promise {
	return &Promise{
		ft: make(Future),
	}
}
func (pm *Promise)Set(v string) {
	pm.ft <- v
}
func (pm *Promise)GetFuture() Future { return pm.ft }
// STOP2 OMIT

// START3 OMIT
func foo(ft Future) {
	v := ft.Get()
	fmt.Println(v)
}
func main() {
	pm := NewPromise()
	ft := pm.GetFuture()
	go foo(ft)
	pm.Set("promise value")
}
// STOP3 OMIT