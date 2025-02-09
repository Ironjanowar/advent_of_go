package main

import (
	"flag"
	"os"
)

type args struct {
	Input string
	Star  int
}

func main() {
	var args args = parseInput()

	data, err := os.ReadFile(args.Input)
	check(err)

	if args.Star == 1 {
		star1(string(data))
	} else {
		star2(string(data))
	}

}

func parseInput() args {
	var input string
	var star int
	flag.StringVar(&input, "input", "", "Input string")
	flag.IntVar(&star, "star", 1, "Star number")
	flag.Parse()

	if input == "" {
		panic("No input file provided")
	}

	return args{Input: input, Star: star}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
