package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	d, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("read %d bytes\n", len(d))

	var (
		input = string(d)
		s     = 0
		e     = 14
		slice string
	)

	for {

		m := map[rune]struct{}{}
		slice = input[s:e]

		for _, r := range slice {
			m[r] = struct{}{}
		}

		if len(m) == 14 {
			break
		}

		s = s + 1
		e = e + 1
	}

	fmt.Println(slice, len(input[:e]))

}
