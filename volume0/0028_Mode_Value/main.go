package main

import (
	"fmt"
	"sort"
)

func main() {
	var max int

	numbers := map[int]int{}
	for {
		var a int
		_, err := fmt.Scan(&a)
		if err != nil {
			break
		}
		numbers[a]++
		if numbers[a] >= max {
			max = numbers[a]
		}
	}

	var modes []int
	for a, b := range numbers {
		if b == max {
			modes = append(modes, a)
		}
	}

	sort.Ints(modes)
	for _, mode := range modes {
		fmt.Println(mode)
	}
}
