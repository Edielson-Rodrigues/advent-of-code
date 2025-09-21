package main

import (
	"advent-of-code/2024/day02/valid_sequences"
	"fmt"
)

func main() {
	fmt.Println("📊 Starting Data Sequence Validation Protocol...")
	fmt.Println("==================================================")

	validSequenceCount := valid_sequences.Calculate()

	fmt.Printf("\n✨ Analysis complete! Found %d valid sequences. ✨\n\n", validSequenceCount)
	fmt.Println("Mission accomplished. Data integrity confirmed. ✅")
	fmt.Println("==================================================")
}
