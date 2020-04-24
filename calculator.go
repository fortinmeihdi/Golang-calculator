package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entrez un opérateur parmie +, -, *, /, % : ")
	scanner.Scan()
	operateur := scanner.Text()

	if operateur == "+" || operateur == "-" || operateur == "*" || operateur == "/" || operateur == "%" {

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Entrez un premier chiffre : ")
		scanner.Scan()
		premier, _ := strconv.Atoi(scanner.Text())

		fmt.Print("Entrez un second chiffre : ")
		scanner.Scan()
		second, _ := strconv.Atoi(scanner.Text())

		switch {
		case operateur == "+":
			println(premier + second)
		case operateur == "-":
			println(premier - second)
		case operateur == "*":
			println(premier * second)
		case operateur == "/":
			println(premier / second)
		case operateur == "%":
			println(premier % second)

		}
	} else {
		fmt.Println("Opérateur inexistant. Veuillez relancer avec un opérateur éxistant")
	}

}
