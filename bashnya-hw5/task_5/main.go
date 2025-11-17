package main

import "fmt"

var (
	ErrIndex = fmt.Errorf("index out of range (0-63)")
	Errmode  = fmt.Errorf("invalid mode: should be 0 or 1")
)

func main() {
	num, i, mode, err := input()
	if err != nil {
		fmt.Print(err)
		return
	}

	res := calc(num, i, mode)
	fmt.Printf("%064b\n%064b\n", num, res)
}

func input() (int64, int, int, error) {
	var num int64
	var i, mode int
	fmt.Println("Input num (int64), bit index(0-63), mode(0 or 1)")
	fmt.Scan(&num, &i, &mode)

	if i < 0 || i >= 64 {
		return 0, 0, 0, ErrIndex
	}

	if mode != 0 && mode != 1 {
		return 0, 0, 0, Errmode
	}

	return num, i, mode, nil
}

func calc(num int64, i int, mode int) int64 {
	if mode == 1 {
		return num | (1 << i)
	}
	return num & ^(1 << i)
}
