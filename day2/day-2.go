package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportStrictlySafe(arr []int) bool {
	increasing := false
	decreasing := false

	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff > 3 || diff < -3 {
			return false
		}
		if diff == 0 {
			return false
		}
		if diff < 0 {
			increasing = true
		}
		if diff > 0 {
			decreasing = true
		}
	}

	return (increasing || decreasing) && !(increasing && decreasing)
}

func isReportSafeWithRemoval(arr []string) bool {
	// Convert string array to int array
	nums := make([]int, len(arr))
	for i, v := range arr {
		nums[i], _ = strconv.Atoi(v)
	}

	// Check if the report is strictly safe
	if isReportStrictlySafe(nums) {
		return true
	}

	// Try removing each level and re-check
	for i := 0; i < len(nums); i++ {
		newArr := append([]int{}, nums[:i]...)
		newArr = append(newArr, nums[i+1:]...)
		fmt.Println("Removed element:", nums[i])
		fmt.Println("Old array:", nums)
		fmt.Println("New array:", newArr)
		fmt.Println("Is strictly safe:", isReportStrictlySafe(newArr))
		if isReportStrictlySafe(newArr) {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeReports := 0

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Fields(line)
		if isReportSafeWithRemoval(arr) {
			safeReports++
		}
	}

	fmt.Println(safeReports)
}
