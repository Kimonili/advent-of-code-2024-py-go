package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ReadInputFile reads number pairs from a file and returns two slices
func ReadInputFile(filename string) ([]int, []int) {
	// here we initialise two empty slices of integers
	list1 := []int{}
	list2 := []int{}

	// Open the input file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split the line into numbers
		numbers := strings.Fields(line) // numbers is now a slice of strings (similar to python list)
		if len(numbers) != 2 {
			continue
		}
		// Convert strings to integers
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	fmt.Printf("total number of pairs: %d\n", len(list1))
	return list1, list2
}

// SortList sorts a slice of integers
func SortList(ls []int) []int {
	// We already know the length of the original slice
	// so we initialize the sorted one with 0s of the same length
	sorted := make([]int, len(ls))
	copy(sorted, ls)
	sort.Ints(sorted)
	return sorted
}

// FindDiff calculates absolute differences between sorted slices
func FindDiff(ls1, ls2 []int) []int {
	if len(ls1) != len(ls2) {
		log.Fatal("Lists must have equal length")
	}

	sortedLs1 := SortList(ls1)
	sortedLs2 := SortList(ls2)

	diff := make([]int, len(ls1))
	for i := range sortedLs1 {
		difference := Abs(sortedLs1[i] - sortedLs2[i])
		diff[i] = difference
	}
	return diff
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CalculateTimesInList(ls []int, num int) int {
	count := 0
	for _, n := range ls {
		if n == num {
			count++
		}
	}
	return count * num
}
