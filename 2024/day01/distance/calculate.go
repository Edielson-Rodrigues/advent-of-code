package distance

import (
	"advent-of-code/pkg/lists"
	"advent-of-code/pkg/numbers"
	"advent-of-code/pkg/sorting"
)

func Calculate(leftList []int, rightList []int) int {
	sorting.MergeSort(leftList)
	sorting.MergeSort(rightList)

	biggerListSize := numbers.GetBiggest(len(leftList), len(rightList))

	totalDistance := 0

	for i := range biggerListSize {
		leftNumber := lists.GetSafe(leftList, i)
		rightNumber := lists.GetSafe(rightList, i)

		distance := numbers.GetBiggest(leftNumber, rightNumber) - numbers.GetSmallest(leftNumber, rightNumber)
		totalDistance += distance
	}

	return totalDistance
}
