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
		a, _ := strconv.Atoi(sc.Text())
		fmt.Println(countPrime(a))
	}
}

func countPrime(a int) int {
	primes := make([]bool, a+1)
	n := 0
	for i := 2; i <= a; i++ {
		if primes[i] == false {
			n++
			for j := i * i; j <= a; j += i {
				primes[j] = true
			}
		}
	}
	return n
}
