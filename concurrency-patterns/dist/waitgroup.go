package main

import (
	"sync"
)

type httpPkg struct{}
func (httpPkg) Get(url string) {}
var http httpPkg

func main() {
// START1 OMIT
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		wg.Add(1).
		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(url)
	}

	wg.Wait()
	// STOP1 OMIT
}
