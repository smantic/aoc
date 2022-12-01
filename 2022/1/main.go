package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var (
		i      int
		ints   []int = make([]int, 30)
		totals []int = make([]int, 0, 100)
	)

	d, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes\n", len(d))

	r := bufio.NewReader(bytes.NewReader(d))
	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			sort.Ints(totals)
			fmt.Println(totals)
			return
		}

		if len(l) == 0 {
			var m int
			for _, i := range ints {
				m = m + i
			}
			totals = append(totals, m)

			i = 0
			zero(ints)
			continue
		}

		ints[i], err = strconv.Atoi(string(l))
		if err != nil {
			panic(err)
		}
		i++
	}

}

func zero(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = 0
	}
}
