package main

import (
	"fmt"
)

func main() {
	var first int
	var second int
	var operator string
	fmt.Println("Print the first number : ")
	fmt.Scan(&first)
	fmt.Println("Print the second number : ")
	fmt.Scan(&second)
	fmt.Println("Enter the operator:  ")
	fmt.Scan(&operator)
	switch operator {
	case "+":
		fmt.Println(calculate(first, second, sumFunk))
	case "-":
		fmt.Println(calculate(first, second, minusFunc))
	case "*":
		fmt.Println(calculate(first, second, multFunk))
	case "/":
		fmt.Println(calculate(first, second, delFunc))
	}
}

var sumFunk = func(x, y int) int {
	return x + y
}
var multFunk = func(x, y int) int {
	return x * y
}
var delFunc = func(x, y int) int {
	return x / y
}
var minusFunc = func(x, y int) int {
	return x - y
}

func calculate(a, b int, action func(x, y int) int) int {
	return action(a, b)
}
