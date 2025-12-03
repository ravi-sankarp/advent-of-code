package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type input struct {
	min int
	max int
}

func parseInput() []input {
	txtBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	strArr := strings.Split(string(txtBytes), ",")

	result := make([]input, len(strArr))

	for i, str := range strArr {
		numRange := strings.Split(str, "-")
		min, _ := strconv.Atoi(strings.Join(strings.Fields(numRange[0]), ""))
		max, _ := strconv.Atoi(strings.Join(strings.Fields(numRange[1]), ""))

		result[i] = input{
			min,
			max,
		}
	}
	return result
}

func part1(input []input) {
	sum := 0
	for _, ele := range input {
		for i := ele.min; i <= ele.max; i++ {
			str := strconv.Itoa(i)
			mid := len(str) / 2

			if str[:mid] == str[mid:] {
				sum += i
			}
		}
	}
	fmt.Println("Part 1 Result is ", sum)
}

func part2(input []input) {
	sum := 0
	for _, ele := range input {
		for i := ele.min; i <= ele.max; i++ {

			str := strconv.Itoa(i)
			mid := len(str) / 2

			if str[:mid] == str[mid:] {
				sum += i
			} else {

				divisors := getDivisors(i)
				matchFound := true
				for _, ele := range divisors {
					matchFound = true
					for i := 0; i < len(str)-ele; i = i + ele {
						if str[i:ele] != str[i+ele:i+ele+i+ele] {
							matchFound = false
							break
						}
					}
					if matchFound {
						break
					}
				}
				if matchFound == true {
					sum += i
				}
			}
		}
	}
	fmt.Println("Part 2 Result is ", sum)
}

func getDivisors(num int) []int {
	numLen := len(strconv.Itoa(num))
	arr := make([]int, 0, numLen/2)
	arr = append(arr, 1)
	for i := 2; i <= numLen/2; i++ {
		if numLen%2 == 0 {
			arr = append(arr, i)
		}
	}
	return arr
}

// link: https://adventofcode.com/2025/day/2
func main() {
	input := parseInput()
	//  part1(input)

	part2(input)
}
