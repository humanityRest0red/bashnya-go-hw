package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"time"
)

//	func foo(cancel context.CancelFunc) {
//		sigCh := make(chan os.Signal, 1)
//		signal.Notify(sigCh, os.Interrupt)
//		go func() {
//			<-sigCh
//			fmt.Println("\n\033[31mInterrupt signal received, shutting down...\033[0m")
//			cancel()
//		}()
//	}
func main() {
	workers, err := getWorkersCount()
	if err != nil {
		fmt.Print(err)
		return
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// foo(cancel)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		fmt.Println("\n\033[31mInterrupt signal received, shutting down...\033[0m")
		cancel()
	}()

	dataCh := make(chan int)

	var wg sync.WaitGroup
	wg.Go(func() {
		sendRandomDataToChannel(ctx, dataCh)
	})

	for w := range workers {
		wg.Go(func() {
			readAndProcess(ctx, dataCh, w+1)
		})
	}
	wg.Wait()

	fmt.Println("All Done... Are you happy now?")

}

func getWorkersCount() (uint, error) {
	var N uint
	flag.UintVar(&N, "N", 0, "Workers count [>0]")
	flag.Parse()

	if N == 0 {
		return 0, fmt.Errorf("usage: %s -N [>0]", os.Args[0])
	}
	return N, nil
}

func sendRandomDataToChannel(ctx context.Context, ch chan int) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- rand.IntN(1000)
		}
	}
}

func readAndProcess(ctx context.Context, ch chan int, wId uint) {
	for {
		select {
		case <-ctx.Done():
			return
		case num, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %v is dead...\n", wId)
				return
			}
			fmt.Printf("Worker %v: %v\n", wId, num)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
