package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getDataPerLine() <-chan []int {
	ch := make(chan []int)

	go func() {
		file, err := os.Open("2024/day02/input.txt")
		if err != nil {
			log.Fatalf("failed to open input: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			dataAsStringArray := strings.Split(line, " ")

			dataAsIntArray := make([]int, len(dataAsStringArray))

			for i := range dataAsStringArray {
				value, err := strconv.Atoi(dataAsStringArray[i])
				if err != nil {
					log.Fatalf("invalid number: %v", err)
				}

				dataAsIntArray[i] = value
			}

			ch <- dataAsIntArray
		}
		close(ch)
	}()

	return ch
}

func getValidationStrategy(num1 int, num2 int) ValidationStrategy {
	if num1 > num2 {
		return DECREASING
	}

	return INCREASING
}

func validateDataSequence(curr int, prev int, validationStrategy ValidationStrategy) bool {
	diff := curr - prev

	switch validationStrategy {
	case INCREASING:
		return curr > prev && diff >= 1 && diff <= 3
	case DECREASING:
		return curr < prev && -diff >= 1 && -diff <= 3
	}

	return true
}

func countValidData() int {
	count := 0

	for data := range getDataPerLine() {
		if len(data) <= 1 {
			count++
			continue
		}

		strategy := getValidationStrategy(data[0], data[1])
		valid := true

		for i := 1; i < len(data); i++ {
			curr := data[i]
			prev := data[i-1]

			if !validateDataSequence(curr, prev, strategy) {
				valid = false
				break
			}
		}

		if valid {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println("📊 Starting Data Sequence Validation Protocol...")
	fmt.Println("==================================================")

	validSequenceCount := countValidData()

	fmt.Printf("\n✨ Analysis complete! Found %d valid sequences. ✨\n\n", validSequenceCount)
	fmt.Println("Mission accomplished. Data integrity confirmed. ✅")
	fmt.Println("==================================================")
}
