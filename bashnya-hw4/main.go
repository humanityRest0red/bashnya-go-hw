package main

import (
	"fmt"
	"myunique"
)

func main() {
	if err := myunique.Run(); err != nil {
		fmt.Println(err)
	}
}
