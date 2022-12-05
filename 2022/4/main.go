package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

	r := bufio.NewReader(bytes.NewReader(d))
	var total int

	for {

		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		l := string(line)

		pair := strings.Split(l, ",")

		right := strings.Split(pair[0], "-")
		left := strings.Split(pair[1], "-")

		var (
			x1, _ = strconv.Atoi(right[0])
			y1, _ = strconv.Atoi(right[1])
			x2, _ = strconv.Atoi(left[0])
			y2, _ = strconv.Atoi(left[1])
		)

		if x1 <= x2 && y1 >= x2 {
			total++
			continue
		}

		if x2 <= x1 && y2 >= x1 {
			total++
			continue
		}

	}
	fmt.Println(total)
}
