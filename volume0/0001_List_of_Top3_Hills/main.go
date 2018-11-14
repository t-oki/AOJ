package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	heights := make([]int, 10)
	for i := 0; i < 10; i++ {
		sc.Scan()
		heights[i], _ = strconv.Atoi(sc.Text())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(heights)))
	for i := 0; i < 3; i++ {
		fmt.Println(heights[i])
	}
}
