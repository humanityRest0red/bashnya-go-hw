package main

import "fmt"

func main() {
	data := []int{1, 3, 0, 6, -5}
	dataChan := sendDataToChannel(data)
	squaredChan := doubleNumbers(dataChan)

	for num := range squaredChan {
		fmt.Println(num)
	}
}

func sendDataToChannel(data []int) chan int {
	ch := make(chan int, len(data))

	go func() {
		defer close(ch)
		for _, num := range data {
			ch <- num
		}
	}()

	return ch
}

func doubleNumbers(inputChan chan int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for num := range inputChan {
			ch <- 2 * num
		}
	}()

	return ch
}
