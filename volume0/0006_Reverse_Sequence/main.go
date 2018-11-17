package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	s := sc.Text()
	l := len(s)
	r := make([]rune, l)
	for i, c := range s {
		r[l-i-1] = c
	}
	fmt.Println(string(r))
}
