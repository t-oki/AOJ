package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	for sc.Scan() {
		d, _ := strconv.Atoi(sc.Text())
		var area int
		for i := 1; i*d < 600; i++ {
			area += d * (d * i) * (d * i)
		}
		fmt.Println(area)
	}
}
