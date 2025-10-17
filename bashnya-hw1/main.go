package main

import "fmt"

func main() {
	left, right, op := Input()
	result := Calc(left, right, op)
	fmt.Printf("%.3f", result)
}

func Input() (float64, float64, string) {
	var left, right float64
	var op string

	fmt.Println("Input left operand:")
	_, err := fmt.Scan(&left)
	for err != nil {
		fmt.Println("Error float number, try again")
		_, err = fmt.Scan(&left)
	}

	fmt.Println("Input operation:")
	_, err = fmt.Scan(&op)
	for err != nil || !IsValidOp(op) {
		if err != nil {
			fmt.Println("Error string input")
		} else {
			fmt.Println("Error opearion, try [+][-][*][/]")
		}
		_, err = fmt.Scan(&op)
	}

	fmt.Println("Input right operand:")
	_, err = fmt.Scan(&right)
	for err != nil {
		fmt.Println(err)
		_, err = fmt.Scan(&right)
	}
	return left, right, op
}

func IsValidOp(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func Calc(left, right float64, op string) float64 {
	var result float64 = 0
	switch op {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	default:
		fmt.Println("Error op")
	}
	return result
}
