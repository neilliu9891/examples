package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func goTimer(ctx context.Context) {
	timer := time.NewTimer(1 * time.Second)
	for {
		rand.Seed(time.Now().UnixNano())
		tt := 1 + rand.Intn(5)
		timer.Reset(time.Duration(tt) * time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("stop timer")
			timer.Stop()
			return
		case <-timer.C:
			fmt.Printf("timer ...%s\n", time.Now().Local().String())
		}
	}
}

func main() {
	fmt.Println("vim-go")
	ctx, cancel := context.WithCancel(context.Background())
	go goTimer(ctx)

	time.Sleep(100 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
