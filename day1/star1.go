package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func star1(testInput string) {
	var lines []string = splitLines(testInput)

	var left, right []int = splitAndSortColumns(lines)

	var distances []int = calculateDistances(left, right)

	var result int = 0

	for _, distance := range distances {
		result += distance
	}

	fmt.Println(result)
}

func splitLines(input string) []string {
	var result []string

	for _, line := range strings.Split(input, "\n") {
		if strings.TrimSpace(line) != "" {
			result = append(result, line)
		}
	}

	return result
}

func splitAndSortColumns(lines []string) ([]int, []int) {
	left, right := split(lines)

	sortColumn(&left)
	sortColumn(&right)

	return left, right
}

func split(lines []string) ([]int, []int) {
	var left, right []int

	for _, line := range lines {
		var splitLine []string
		for _, elem := range strings.Split(line, " ") {
			trimmed := strings.TrimSpace(elem)
			if trimmed != "" {
				splitLine = append(splitLine, trimmed)
			}
		}

		leftInt, err := strconv.Atoi(splitLine[0])
		check(err)

		rightInt, err := strconv.Atoi(splitLine[1])
		check(err)

		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	return left, right
}

func sortColumn(array *[]int) {
	sort.Ints(*array)
}

func calculateDistances(left []int, right []int) []int {
	var distances []int

	if len(left) != len(right) {
		panic("Left and right columns are not the same length")
	}

	for i := 0; i < len(left); i++ {
		distances = append(distances, abs(left[i]-right[i]))
	}

	return distances
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
