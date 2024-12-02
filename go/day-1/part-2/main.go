package main

import (
	"day1/utils"
	"fmt"
)

func main() {
	list1, list2 := utils.ReadInputFile("../input")

	// Initialise a slice of ints with the same length as the first list
	similarityList := make([]int, len(list1))

	// Iterate over the second list
	for i, num1 := range list1 {

		timesInList := utils.CalculateTimesInList(list2, num1)

		simScore := timesInList * num1

		similarityList[i] = simScore
	}

	// Calculate sum of similarity scores
	sum := 0
	for _, score := range similarityList {
		sum += score
	}

	fmt.Printf("total similarity score: %d\n", sum)
}
