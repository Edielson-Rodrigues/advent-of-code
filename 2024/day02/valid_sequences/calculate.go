package valid_sequences

import (
	"bufio"
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

func isValidReport(data []int) bool {
	if len(data) <= 1 {
		return true
	}

	strategy := getValidationStrategy(data[0], data[1])

	for i := 1; i < len(data); i++ {
		if !validateDataSequence(data[i], data[i-1], strategy) {
			return false
		}
	}

	return true
}

func Calculate() int {
	count := 0

	for data := range getDataPerLine() {
		isSafe := isValidReport(data)

		if isSafe {
			count++
			continue
		}

		for i := range data {
			candidate := append([]int{}, data[:i]...)
			candidate = append(candidate, data[i+1:]...)

			isSafe = isValidReport(candidate)
			if isSafe {
				count++
				break
			}
		}
	}

	return count
}
