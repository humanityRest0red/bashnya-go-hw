package main

import (
	"fmt"
	"sync"
)

func main() {
	var data = []int{2, 4, 6, 8, 10}
	sum := 0

	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, num := range data {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			val := x * x
			mu.Lock()
			sum += val
			mu.Unlock()
		}(num)
	}
	wg.Wait()

	fmt.Println(sum)
}
