package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	dataset := []int{a, b, c, d, e}

	sort.Sort(sort.Reverse(sort.IntSlice(dataset)))
	fmt.Printf("%d %d %d %d %d\n", dataset[0], dataset[1], dataset[2], dataset[3], dataset[4])
}
