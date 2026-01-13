package concurrent_map_test

import (
	"concurrent_map"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func BenchmarkMapAdd(b *testing.B) {
	const n = 1_000_000

	for _, workers := range []int{1, 2, runtime.NumCPU(), 5 * runtime.NumCPU()} {
		bName := "concurent map/w=" + strconv.Itoa(workers)
		b.Run(bName, func(b *testing.B) {
			for b.Loop() {
				cm := concurrent_map.New[string, int]()

				var wg sync.WaitGroup
				wg.Add(workers)

				for w := range workers {
					go func(seed int64) {
						defer wg.Done()
						r := rand.New(rand.NewSource(time.Now().UnixNano() + seed))
						for range n {
							letter := charset[r.Intn(len(charset))]
							cm.Add(string(letter), r.Intn(1_000_000))
						}
					}(int64(w))
				}
				wg.Wait()
			}
		})
	}

	b.Run("map", func(b *testing.B) {
		for b.Loop() {
			m := make(map[string]int)
			for range n {
				letter := charset[rand.Intn(len(charset))]
				m[string(letter)] = rand.Intn(1_000_000)
			}
		}
	})
}

func BenchmarkMapGet(b *testing.B) {
	const n = 1_000_000
	cm := concurrent_map.NewFromMap(*GenMapN(n))
	m := *GenMapN(n)

	for _, workers := range []int{1, 2, runtime.NumCPU(), 2 * runtime.NumCPU(), 5 * runtime.NumCPU()} {
		bName := "concurent map/w=" + strconv.Itoa(workers)
		b.Run(bName, func(b *testing.B) {
			for b.Loop() {
				var wg sync.WaitGroup
				wg.Add(workers)

				for range workers {
					go func() {
						defer wg.Done()
						for _, letter := range charset {
							cm.Get(string(letter))
						}
					}()
				}
				wg.Wait()
			}
		})
	}

	b.Run("map", func(b *testing.B) {
		for b.Loop() {
			for _, letter := range charset {
				_ = m[string(letter)]
			}
		}
	})
}

func GenMapN(n int) *map[string]int {
	m := make(map[string]int, n)

	for range n {
		letter := charset[rand.Intn(len(charset))]
		m[string(letter)] = rand.Intn(1_000_000)
	}

	return &m
}
