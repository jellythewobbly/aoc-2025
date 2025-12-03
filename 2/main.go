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

	p2 := part2(&input)
	fmt.Printf("Part 2: %v\n", p2)
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
			if check > highInt {
				break
			}

			if check >= lowInt {
				sum += check
			}

			digitInt++
		}
	}

	return sum
}

func part2(input *string) int64 {
	sum := int64(0)

	for idRange := range strings.SplitSeq(*input, ",") {
		split := strings.Split(idRange, "-")
		low, high := split[0], split[1]
		lowInt, _ := strconv.ParseInt(low, 10, 64)
		highInt, _ := strconv.ParseInt(high, 10, 64)

		minLen := len(low)
		maxLen := len(high)

		seen := make(map[int64]bool)

		for totalLen := minLen; totalLen <= maxLen; totalLen++ {
			currentInt := lowInt
			if totalLen > minLen {
				currentInt = int64(math.Pow10(totalLen - 1))
			}

			currentStr := strconv.FormatInt(currentInt, 10)

			for patternLen := 1; patternLen <= totalLen/2; patternLen++ {
				if totalLen%patternLen != 0 {
					continue
				}

				repeatCount := totalLen / patternLen

				digits := currentStr[:patternLen]
				patternInt, _ := strconv.ParseInt(digits, 10, 64)
				maxPattern := int64(math.Pow10(patternLen)) - 1

				for patternInt := patternInt; patternInt <= maxPattern; patternInt++ {
					patternStr := strconv.FormatInt(patternInt, 10)
					repeated := strings.Repeat(patternStr, repeatCount)
					check, _ := strconv.ParseInt(repeated, 10, 64)

					if check > highInt {
						break
					}

					if check >= lowInt && !seen[check] {
						seen[check] = true
						sum += check
					}
				}
			}
		}
	}

	return sum
}

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
