package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums := part1()
	fmt.Println(nums)

	nums2 := part2()
	fmt.Println(nums2)
}

func part1() int {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		intNums, err := convertToIntSlice(nums)
		if err != nil {
			panic(err)
		}

		increasing := intNums[0] <= intNums[1]
		changed := false

		for i := 1; i < len(intNums); i++ {
			diff := intNums[i] - intNums[i-1]

			if (diff < 0 && increasing) || (diff > 0 && !increasing) || abs(diff) > 3 || diff == 0 {
				changed = true
				break
			}
		}

		if !changed {
			safe++
		}
	}
	return safe
}

func part2() int {
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	scanner := bufio.NewScanner(fi)
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		intNums, err := convertToIntSlice(nums)
		if err != nil {
			panic(err)
		}

		increasing := intNums[0] <= intNums[1]
		changed := 0

		for i := 1; i < len(intNums); i++ {
			diff := intNums[i] - intNums[i-1]

			if (diff < 0 && increasing) || (diff > 0 && !increasing) || abs(diff) > 3 || diff == 0 {
				changed++
			}
		}

		if changed <= 1 {
			safe++
		}
	}
	return safe
}

func convertToIntSlice(nums []string) ([]int, error) {
	intNums := make([]int, len(nums))
	for i, num := range nums {
		var err error
		intNums[i], err = strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
	}
	return intNums, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
