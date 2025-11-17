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

func main() {
	var (
		workers = mustWorkersCount()
		ctx     = contextWithSignalInterrupt()
		dataCh  = make(chan int)
		wg      sync.WaitGroup
	)

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

func mustWorkersCount() uint {
	var N uint
	flag.UintVar(&N, "N", 0, "Workers count [>0]")
	flag.Parse()

	if N == 0 {
		panic(fmt.Sprintf("usage: %s -N [>0]", os.Args[0]))
	}

	return N
}

func contextWithSignalInterrupt() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		<-sigCh
		fmt.Println("\n\033[31mInterrupt signal received, shutting down...\033[0m")
		cancel()
	}()

	return ctx
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
