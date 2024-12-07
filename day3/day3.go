package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readFileIntoString(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputText := ""

	for scanner.Scan() {
		for _, char := range scanner.Text() {
			inputText += string(char)
		}
	}
	return inputText
}

func parseMulOperations(inputText, regex string) [][]string {
	re := regexp.MustCompile(regex)

	return re.FindAllStringSubmatch(inputText, -1)
}

// Parses the input string to match the mul operation, then multiplies all matches to get total
func part1() {
	inputString := readFileIntoString("input.txt")

	matches := parseMulOperations(inputString, `mul\((\d{1,3}),(\d{1,3})\)`)
	total := 0
	for i := 0; i < len(matches); i++ {
		num1, _ := strconv.Atoi(matches[i][1])
		num2, _ := strconv.Atoi(matches[i][2])
		total += num1 * num2
	}

	fmt.Println("Answer for part1: ", total)
}

// Same as part one except we add the dont/do operations and track if we should multiply
func part2() {
	inputString := readFileIntoString("input.txt")

	matches := parseMulOperations(inputString, `don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	shouldMultiply := true

	total := 0
	for i := 0; i < len(matches); i++ {
		if matches[i][0] == "do()" {
			shouldMultiply = true
			continue
		}
		if matches[i][0] == "don't()" {
			shouldMultiply = false
			continue
		}

		if shouldMultiply {
			num1, _ := strconv.Atoi(matches[i][1])
			num2, _ := strconv.Atoi(matches[i][2])
			total += num1 * num2

		}
	}

	fmt.Println("Answer for part2: ", total)
}

func main() {
	part1()
	part2()
}
