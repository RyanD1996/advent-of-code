package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Sorts both lists and then calculates the difference bwtween elect equivalent element in both lists
func accumulatedDiff(list1, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)
	var accumulatedDiff int = 0
	for i := 0; i < len(list1); i++ {
		num1 := list1[i]
		num2 := list2[i]

		if num1 < num2 {
			accumulatedDiff = accumulatedDiff + (num2 - num1)
		} else {
			accumulatedDiff = accumulatedDiff + (num1 - num2)
		}
	}
	return accumulatedDiff
}

func createFrequencyList(list []int) map[int]int {
	m := make(map[int]int)

	for i := 0; i < len(list); i++ {
		element := list[i]
		m[element] = m[element] + 1
	}

	return m
}

func calculateSimilarityScore(list1, list2 []int) int {
	freqList2 := createFrequencyList(list2)

	totalSimilarityScore := 0

	for i := 0; i < len(list1); i++ {
		totalSimilarityScore = totalSimilarityScore + (list1[i] * freqList2[list1[i]])
	}

	return totalSimilarityScore

}

func main() {
	file, err := os.Open("input.txt")
	var list1, list2 []int = make([]int, 0), make([]int, 0)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read file and parse integers
	for scanner.Scan() {
		var a, b int
		line := scanner.Text()

		_, err := fmt.Sscanf(line, "%d %d", &a, &b)
		if err != nil {
			fmt.Println("Error parsing file: ", err)
			panic(err)
		}

		list1 = append(list1, a)
		list2 = append(list2, b)
	}

	// Sort lists
	if len(list1) != len(list2) {
		fmt.Println("Error: Lists are not the same length")
		return
	}

	// Part one answer
	accumulatedDiff := accumulatedDiff(list1, list2)
	fmt.Println("Accumulated difference: ", accumulatedDiff)

	// Part two answer
	similarityScore := calculateSimilarityScore(list1, list2)
	fmt.Println("Similarity score: ", similarityScore)

}
