package main

import (
	"day1/utils"
	"fmt"
)

func main() {
	list1, list2 := utils.ReadInputFile("../input")
	diff := utils.FindDiff(list1, list2)

	// Calculate sum of differences
	// In go it is common to use a range to iterate over a slice
	// when we set the first variable to `_` we don't use it (similar to python)
	sum := 0
	for _, d := range diff {
		sum += d
	}

	fmt.Printf("total diff: %d\n", sum)
}
