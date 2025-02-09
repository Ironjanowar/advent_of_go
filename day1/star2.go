package main

import (
	"fmt"
)

func star2(testInput string) {
	var lines []string = splitLines(testInput)

	var left, right []int = split(lines)

	var counts map[int]int = countApparitions(right)

	var result int = multiplyByCountAndSum(left, counts)

	fmt.Println(result)
}

func countApparitions(list []int) map[int]int {
	var counts map[int]int = make(map[int]int)

	for _, elem := range list {
		if _, ok := counts[elem]; ok {
			counts[elem]++
		} else {
			counts[elem] = 1
		}
	}

	return counts
}

func multiplyByCountAndSum(list []int, counts map[int]int) int {
	var result int = 0

	for _, elem := range list {
		if count, ok := counts[elem]; ok {
			result += elem * count
		}
	}

	return result
}
