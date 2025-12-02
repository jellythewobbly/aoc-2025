package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	p1 := part1(&input)
	fmt.Printf("Part 1: %v\n", p1)

	// p2 := part2(&input)
	// fmt.Printf("Part 2: %v\n", p2)
}

func part1(input *string) int64 {
	sum := int64(0)

	for idRange := range strings.SplitSeq(*input, ",") {
		split := strings.Split(idRange, "-")
		low, high := split[0], split[1]
		lowInt, _ := strconv.ParseInt(low, 10, 64)
		highInt, _ := strconv.ParseInt(high, 10, 64)

		if len(low)%2 == 1 && len(low) == len(high) {
			continue
		}

		currentInt := lowInt
		if len(low)%2 == 1 {
			currentInt = int64(math.Pow10(len(low)))
		}

		if currentInt > highInt {
			continue
		}

		id := strconv.FormatInt(currentInt, 10)
		digits := id[:len(id)/2]
		digitInt, _ := strconv.ParseInt(digits, 10, 64)

		for {
			check, _ := strconv.ParseInt(fmt.Sprintf("%v%v", digitInt, digitInt), 10, 64)
			if check >= lowInt && check <= highInt {
				sum += check
			}

			if check > highInt {
				break
			}

			digitInt++
		}

		fmt.Println(lowInt, highInt)
	}

	return sum
}

// func part2(input *string) int64 {
// }

func getInput() string {
	fileReader, err := os.Open("./input")
	if err != nil {
		fmt.Println("error in reading file")
		return ""
	}

	defer fileReader.Close()

	scanner := bufio.NewScanner(fileReader)

	var res string

	scanner.Scan()
	res = scanner.Text()

	return res
}
