package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	dataset, _ := strconv.Atoi(sc.Text())
	var firstdata *big.Int

	for i := 0; i < dataset*2; i++ {
		n := new(big.Int)
		sc.Scan()
		data, _ := n.SetString(sc.Text(), 10)

		if i%2 == 0 {
			firstdata = data
			continue
		}

		result := new(big.Int)
		if result := result.Add(data, firstdata); len(result.String()) > 80 {
			fmt.Println("overflow")
			continue
		} else {
			fmt.Println(result)
		}
	}
}
