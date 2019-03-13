package main

import (
	"context"
	"fmt"
	"time"
)

// START0 OMIT
func loop(ctx context.Context) {
	for i := 0; ; i++ {
		select {
		case <- ctx.Done():
			fmt.Println("done")
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	go loop(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(5 * time.Second)
}
// STOP0 OMIT