package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	var ok bool
	for i := 0; i < len(s)-1; i++ {
		copied := CloneSString(s)
		copied[i], copied[i+1] = copied[i+1], copied[i]
		if IsSortedSString(copied) {
			ok = true
			break
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func IsSortedSString(a []string) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func CloneSString(a []string) []string {
	cloned := make([]string, len(a))
	copy(cloned, a)
	return cloned
}

func CloneSInt(a []int) []int {
	cloned := make([]int, len(a))
	copy(cloned, a)
	return cloned
}
