package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator <operation> <operand1> <operand2>")
		return
	}

	operation := os.Args[1]
	operand1, err1 := strconv.ParseFloat(os.Args[2], 64)
	operand2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input. Mohon untuk masukan operation yang tepat.")
		return
	}

	var result float64
	switch operation {
	case "add":
		result = operand1 + operand2
	case "sub":
		result = operand1 - operand2
	case "mul":
		result = operand1 * operand2
	case "div":
		result = operand1 / operand2
	default:
		fmt.Println("Invalid operation.")
		return
	}

	fmt.Println("Result: ", result)

}