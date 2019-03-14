package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

// START1 OMIT
func Map(ss []string, f func(string) string) []string {
	oss := make([]string, len(ss))
	n := len(ss)
	nc := runtime.NumCPU()
	if nc > n {
		nc = n
	}
	var wg sync.WaitGroup
	for i := 0; i < nc; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for k := 0 + j; k < n; k += nc {
				oss[k] = f(ss[k])
			}
		}()
	}
	wg.Wait()
	return oss
}
// STOP1 OMIT

// START2 OMIT
func Filter(ss []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range ss {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
// STOP2 OMIT