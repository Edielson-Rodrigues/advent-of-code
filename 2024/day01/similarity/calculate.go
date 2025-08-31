package similarity

import (
	"slices"
)

func countOccurrences(target int, list []int) int {
	count := 0

	for i := range len(list) {
		if list[i] == target {
			count++
		}
	}

	return count
}

func Calculate(leftList []int, rightList []int) int {
	verifiedNumber := []int{}
	totalSimilarity := 0

	for i := range leftList {
		leftValue := leftList[i]

		if slices.Contains(verifiedNumber, leftValue) {
			continue
		}

		occurrences := countOccurrences(leftValue, rightList)
		totalSimilarity += leftValue * occurrences

		verifiedNumber = append(verifiedNumber, leftValue)
	}

	return totalSimilarity
}
