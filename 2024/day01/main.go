package main

import (
	"advent-of-code/2024/day01/distance"
	"advent-of-code/2024/day01/similarity"
	"fmt"
)

func main() {
	fmt.Println("📜 Reconciling the Historians' Location ID Lists...")
	fmt.Println("==================================================")

	totalDistance := distance.Calculate(LeftList, RightList)
	fmt.Printf("\n✨ The total calculated distance is: %d ✨\n\n", totalDistance)
	fmt.Println("==================================================")

	totalSimilarity := similarity.Calculate(LeftList, RightList)
	fmt.Printf("\n🔍 The total similarity score is: %d 🔍\n\n", totalSimilarity)
	fmt.Println("==================================================")
}
