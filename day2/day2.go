package main

import (
	"advent_of_go/helper"
)

func main() {
	args, data := helper.ParseData()

	if args.Star == 1 {
		star1(data)
	} else {
		star2(data)
	}
}
