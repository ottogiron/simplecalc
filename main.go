package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var reader *bufio.Reader

	reader = bufio.NewReader(os.Stdin)

	fmt.Println("Enter operation 1) sum 2) rest 3) multiplication 4) division")

	var operation string
	operation, _ = reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	fmt.Println("Operation is:", operation)

	var operand1 float64
	var operand1Str string
	fmt.Println("Operand 1:")
	operand1Str, _ = reader.ReadString('\n')
	operand1Str = strings.TrimSpace(operand1Str)
	operand1, _ = strconv.ParseFloat(operand1Str, 64)

	var operand2 float64
	var operand2Str string
	fmt.Println("Operand 2:")
	operand2Str, _ = reader.ReadString('\n')
	operand2Str = strings.TrimSpace(operand2Str)
	operand2, _ = strconv.ParseFloat(operand2Str, 64)

	if operation == "sum" {
		fmt.Println("Sum result is:", sum(operand1, operand2))
	}

}

func sum(op float64, op2 float64) float64 {
	return op + op2
}
