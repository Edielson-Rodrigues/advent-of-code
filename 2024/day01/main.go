package main

import (
	"advent-of-code/pkg/lists"
	"advent-of-code/pkg/numbers"
	"advent-of-code/pkg/sorting"
	"fmt"
)

func getTotalDistance(leftList []int, rightList []int) int {
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

func main() {
	fmt.Println("📜 Reconciling the Historians' Location ID Lists...")
	fmt.Println("==================================================")

	sorting.MergeSort(LeftList)
	sorting.MergeSort(RightList)

	totalDistance := getTotalDistance(LeftList, RightList)

	fmt.Printf("\n✨ The total calculated distance is: %d ✨\n\n", totalDistance)
	fmt.Println("Mission complete! The lists have been reconciled. ✅")
	fmt.Println("==================================================")
}
