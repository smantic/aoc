package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// a, x -> rock		-> 1
// b, y -> paper    -> 2
// c, z -> scissors -> 3
// 0 -> loss
// 3 -> draw
// 6 -> win

// x -> lose
// y -> draw
// z -> win

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

	var p1, p2 int

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		switch string(line) {
		case "A X":
			p1 = p1 + 7
			p2 = p2 + 3
		case "A Y":
			p1 = p1 + 4
			p2 = p2 + 4
		case "A Z":
			p1 = p1 + 1
			p2 = p2 + 8
		case "B X":
			p1 = p1 + 8
			p2 = p2 + 1
		case "B Y":
			p1 = p1 + 5
			p2 = p2 + 5
		case "B Z":
			p1 = p1 + 2
			p2 = p2 + 9
		case "C X":
			p1 = p1 + 9
			p2 = p2 + 2
		case "C Y":
			p1 = p1 + 6
			p2 = p2 + 6
		case "C Z":
			p1 = p1 + 3
			p2 = p2 + 7
		}
	}

	fmt.Println(p1, p2)

}
