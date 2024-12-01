package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func readInput() ([]int, []int) {
	fi, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	left_list := make([]int, 0, 50)
	right_list := make([]int, 0, 50)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		if len(nums) != 2 {
			log.Fatalf("expected 2 numbers, got %d", len(nums))
		}
		left, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("failed to parse left number: %v", err)
		}
		right, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("failed to parse right number: %v", err)
		}
		left_list = append(left_list, left)
		right_list = append(right_list, right)
	}
	return left_list, right_list
}

func part1() int {
	left_list, right_list := readInput()

	sort.Ints(left_list)
	sort.Ints(right_list)

	res := 0
	for i := range left_list {
		diff := left_list[i] - right_list[i]
		if diff < 0 {
			diff = -diff
		}
		res += diff
	}
	return res
}

func part2() int {
	left_list, right_list := readInput()

	freq_dict := make(map[int]int)
	for _, num := range right_list {
		freq_dict[num]++
	}

	res := 0
	for _, num := range left_list {
		res += (freq_dict[num] * num)
	}
	return res
}
