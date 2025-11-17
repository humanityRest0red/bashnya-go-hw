package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
)

func main() {
	var workers = 2
	var wg sync.WaitGroup
	wg.Add(workers)

	for range workers {
		go func() {
			defer wg.Done()
			ch1 := sendRandomDataToChannel()
			ch2 := readFromChannel(ch1)
			outputChannel(ch2)
		}()
	}
	wg.Wait()

	fmt.Println("done")
}

func sendRandomDataToChannel() chan int {
	ch := make(chan int)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		fmt.Println("\nПолучен сигнал прерывания")
		// cancel()
	}()

	go func() {
		defer close(ch)
		for {
			ch <- rand.IntN(1000)
		}
	}()

	return ch
}

func readFromChannel(data_chan chan int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for num := range data_chan {
			ch <- num
		}
	}()

	return ch
}

func outputChannel(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}
