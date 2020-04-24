package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var operator string

findOperator:
	fmt.Print("Please choose a operator between +, -, *, / or  % : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operator = scanner.Text()

	scanner = bufio.NewScanner(os.Stdin)

	fmt.Print("Enter first number : ")
	scanner.Scan()
	number1, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter second number : ")
	scanner.Scan()
	number2, _ := strconv.Atoi(scanner.Text())

	switch operator {
	case "+":
		println(number1 + number2)
	case "-":
		println(number1 - number2)
	case "*":
		println(number1 * number2)
	case "/":
		println(number1 / number2)
	case "%":
		println(number1 % number2)
	default:
		fmt.Println("Invalid operator select :", operator, ". Please retry.")
		goto findOperator
	}
}
