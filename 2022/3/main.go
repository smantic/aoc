package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

var priority map[rune]int = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

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

	r := bufio.NewReader(bytes.NewReader(d))

	shared := make([]rune, 100)
	i := 0

	for {
		var lines [3]string
		var err error

		for i := 0; i < 3; i++ {
			var l []byte
			l, _, err = r.ReadLine()
			lines[i] = string(l)
		}

		if err == io.EOF {
			break
		}

		counts := make(map[rune]int)

		for _, l := range lines {
			var items = make(map[rune]struct{})
			for _, c := range l {
				items[rune(c)] = struct{}{}
			}

			for c, _ := range items {
				counts[c] += 1
			}

		}

		for c, count := range counts {
			if count == 3 {
				shared[i] = c
			}
		}
		i = i + 1
	}

	var total int
	for _, r := range shared {
		total += priority[r]
	}

	fmt.Println(total)
}
