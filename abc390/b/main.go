package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// N
	// A1 A2 ... An
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	nStr := sc.Text()
	n, err := strconv.Atoi(nStr)
	if err != nil {
		panic(err)
	}
	sc.Scan()
	asStr := strings.Split(sc.Text(), " ")
	as := make([]int, n)
	for i, aStr := range asStr {
		a, err := strconv.Atoi(aStr)
		if err != nil {
			panic(err)
		}
		as[i] = a
	}

	ok := true
	// If len(as) is 2, it's "Yes". Skip this loop.
	for i := 0; i < len(as)-2; i++ {
		if as[i]*as[i+2] != as[i+1]*as[i+1] {
			ok = false
			break
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
