package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go DoFor(&wg)
	}

	wg.Wait()
}
