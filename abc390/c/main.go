package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// H W
	// S1
	// S2
	// ...
	// SH
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	hwStr := strings.Split(sc.Text(), " ")
	h, err := strconv.Atoi(hwStr[0])
	if err != nil {
		panic(err)
	}
	w, err := strconv.Atoi(hwStr[1])
	if err != nil {
		panic(err)
	}
	s := make([]string, 0, h)
	for i := 0; i < h; i++ {
		sc.Scan()
		s = append(s, sc.Text())
	}

	var ti int
	breaked := false
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				ti = i
				breaked = true
				break
			}
		}
		// goto使った方がいいんだろうね
		if breaked {
			break
		}
	}

	var bi int
	breaked = false
	for i := h - 1; i >= 0; i-- {
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				bi = i
				breaked = true
				break
			}
		}
		if breaked {
			break
		}
	}

	var lj int
	breaked = false
	for j := 0; j < w; j++ {
		for i := 0; i < h; i++ {
			if s[i][j] == '#' {
				lj = j
				breaked = true
				break
			}
		}
		if breaked {
			break
		}
	}

	var rj int
	breaked = false
	for j := w - 1; j >= 0; j-- {
		for i := 0; i < h; i++ {
			if s[i][j] == '#' {
				rj = j
				breaked = true
				break
			}
		}
		if breaked {
			break
		}
	}

	ok := true
	breaked = false
	for i := ti; i <= bi; i++ {
		for j := lj; j <= rj; j++ {
			if s[i][j] == '.' {
				ok = false
				breaked = true
				break
			}
		}
		if breaked {
			break
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
