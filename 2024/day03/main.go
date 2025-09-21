package main

import (
	"advent-of-code/2024/day03/input_handlers"
	"advent-of-code/2024/day03/instructions"
	"fmt"
)

func main() {
	fmt.Println("⚙️  Initiating Corrupted Memory Analysis...")
	fmt.Println("==================================================")

	cleanedInstructions := input_handlers.GetCleanedInstructions()
	instructionsResult := instructions.CalculateResult(cleanedInstructions)

	fmt.Printf("\n✅ Memory Scan Complete! Sum of results: %d. ✅\n\n", instructionsResult)
	fmt.Println("Program logic successfully decoded. System stable. 🚀")
	fmt.Println("==================================================")
}
