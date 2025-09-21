package input_handlers

import (
	"os"
	"regexp"
)

func GetCleanedInstructions() []string {
	file, err := os.ReadFile("2024/day03/input_handlers/input.txt")
	if err != nil {
	}

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	cleanedInstructions := re.FindAllString(string(file), -1)

	return cleanedInstructions
}
