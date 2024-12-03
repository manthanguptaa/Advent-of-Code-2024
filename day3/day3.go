package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part1()
	part2()
}

func readInput() []string {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func part1() {
	lines := readInput()
	re := regexp.MustCompile(`mul\(\s*(\d+)\s*,\s*(\d+)\s*\)`)
	res := 0
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			res += a * b
		}
	}
	fmt.Println(res)
}

func part2() {
	lines := readInput()
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	res := 0
	switching := true

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				switching = true
			} else if match[0] == "don't()" {
				switching = false
			} else if len(match) == 3 && switching {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				res += a * b
			}
		}
	}
	fmt.Println(res)
}
