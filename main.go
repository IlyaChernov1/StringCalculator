package main

import (
	"StringCalculator/calculator"
	"StringCalculator/input"
	"fmt"
	"strings"
)

func main() {
	inputStr := input.GetInput()

	result := calculator.Calculate(inputStr)
	result = strings.TrimSpace(result)
	fmt.Printf("Результат: %s\n", result)
}
