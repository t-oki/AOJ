package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	l := sc.Text()
	words := strings.Split(l, " ")

	frequency := map[string]int{}
	long := map[string]int{}
	var maxFrequency, longest int
	var frequentWord, longestWord string
	for _, w := range words {
		frequency[w]++
		long[w] = len(w)
		if frequency[w] > maxFrequency {
			maxFrequency = frequency[w]
			frequentWord = w
		}
		if len(w) > longest {
			longest = long[w]
			longestWord = w
		}
	}

	fmt.Println(frequentWord + " " + longestWord)
}
