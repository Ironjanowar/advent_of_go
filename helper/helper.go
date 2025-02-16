package helper

import (
	"flag"
	"os"
)

type args struct {
	Input string
	Star  int
}

func ParseData() (args, string) {
	var args args = parseInput()

	data, err := os.ReadFile(args.Input)
	check(err)

	return args, string(data)
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
