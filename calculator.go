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

	switch operator {
	case "+":
		fmt.Print("Enter first number : ")
		scanner.Scan()
		number1, _ := strconv.ParseFloat(scanner.Text(), 64)

		fmt.Print("Enter second number : ")
		scanner.Scan()
		number2, _ := strconv.ParseFloat(scanner.Text(), 64)

		result := number1 + number2
		fmt.Printf("%f\n", result)
	case "-":
		fmt.Print("Enter first number : ")
		scanner.Scan()
		number1, _ := strconv.ParseFloat(scanner.Text(), 64)

		fmt.Print("Enter second number : ")
		scanner.Scan()
		number2, _ := strconv.ParseFloat(scanner.Text(), 64)

		result := number1 - number2
		fmt.Printf("%f\n", result)
	case "*":
		fmt.Print("Enter first number : ")
		scanner.Scan()
		number1, _ := strconv.ParseFloat(scanner.Text(), 64)

		fmt.Print("Enter second number : ")
		scanner.Scan()
		number2, _ := strconv.ParseFloat(scanner.Text(), 64)

		result := number1 * number2
		fmt.Printf("%f\n", result)
	case "/":
		fmt.Print("Enter first number : ")
		scanner.Scan()
		number1, _ := strconv.ParseFloat(scanner.Text(), 64)

		fmt.Print("Enter second number : ")
		scanner.Scan()
		number2, _ := strconv.ParseFloat(scanner.Text(), 64)

		result := number1 / number2
		fmt.Printf("%f\n", result)
	case "%":
		fmt.Print("Reminder : Modulo can only be used with whole number !")

		fmt.Print("Enter first number : ")
		scanner.Scan()
		number1, _ := strconv.Atoi(scanner.Text())

		fmt.Print("Enter second number : ")
		scanner.Scan()
		number2, _ := strconv.Atoi(scanner.Text())

		println(number1 % number2)
	default:
		fmt.Println("Invalid operator select :", operator, ". Please retry.")
		goto findOperator
	}
}
