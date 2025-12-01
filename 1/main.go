package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := getInput()
	p1 := part1(&input)
	fmt.Printf("Part 1: %v\n", p1)

	p2 := part2(&input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(input *[]string) int64 {
	var count int64
	current := 50

	for _, row := range *input {
		sign := 1
		if row[0] == 'L' {
			sign = -1
		}

		clicks, _ := strconv.Atoi(row[1:])
		clicks %= 100

		current += sign * clicks

		if current >= 100 {
			current -= 100
		} else if current < 0 {
			current += 100
		}

		if current == 0 {
			count++
		}
	}

	return count
}

func part2(input *[]string) int64 {
	var count int64
	current := 50

	for _, row := range *input {
		clicks, _ := strconv.Atoi(row[1:])
		passes := int64(clicks / 100)
		clicks %= 100

		if row[0] == 'L' {
			if current != 0 && current <= clicks {
				passes++
			}

			current -= clicks % 100
			if current < 0 {
				current += 100
			}
		} else {
			if clicks+current >= 100 {
				passes++
			}

			current += clicks % 100
			if current >= 100 {
				current -= 100
			}
		}

		count += passes
	}

	return count
}

func getInput() []string {
	fileReader, err := os.Open("./input")
	if err != nil {
		fmt.Println("error in reading file")
		return []string{}
	}

	defer fileReader.Close()

	scanner := bufio.NewScanner(fileReader)

	var res []string

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}
