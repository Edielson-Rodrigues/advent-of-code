package instructions

import (
	"log"
	"strconv"
	"strings"
)

func extractNumberFromExpression(expression string) (int, int) {
	cleanedExpression := strings.TrimSuffix(strings.TrimPrefix(expression, "mul("), ")")

	numbers := strings.Split(cleanedExpression, ",")
	if len(numbers) != 2 {
		log.Fatalf("invalid data: ", numbers)
	}

	num1, err := strconv.Atoi(numbers[0])
	if err != nil {
		log.Fatalf("failed to open input: %v", err)
	}

	num2, err := strconv.Atoi(numbers[1])
	if err != nil {
		log.Fatalf("failed to open input: %v", err)
	}

	return num1, num2
}

func CalculateResult(data []string) int {
	sum := 0
	expressionEnable := true

	for _, item := range data {
		switch item {
		case "do()":
			expressionEnable = true
		case "don't()":
			expressionEnable = false
		default:
			if expressionEnable {
				num1, num2 := extractNumberFromExpression(item)

				sum += num1 * num2
			}
		}
	}

	return sum
}
