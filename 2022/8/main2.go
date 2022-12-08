package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
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

	const size int = 99
	var (
		line   int = 0
		trees      = make([][]int, size)
		scores     = make([][]int, size)
	)

	for i, _ := range trees {
		trees[i] = make([]int, size)
		scores[i] = make([]int, size)

		for j, _ := range scores[i] {
			scores[i][j] = 1
		}
	}

	// get tree hieghts
	for {

		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		for j, c := range string(l) {
			height, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			trees[line][j] = height

		}
		line++
	}

	// left to right
	for i := 0; i < len(trees); i++ {
		for k := 0; k < len(trees[i]); k++ {
			for j := k + 1; j < len(trees[i]); j++ {
				if trees[i][j] >= trees[i][k] || j == len(trees[i])-1 {
					scores[i][k] = scores[i][k] * (j - k)
					break
				}
			}
		}
	}

	// right to left
	for i := 0; i < len(trees); i++ {
		for k := len(trees[i]) - 1; k >= 0; k-- {
			for j := k - 1; j >= 0; j-- {
				if trees[i][j] >= trees[i][k] || j == 0 {
					scores[i][k] = scores[i][k] * (k - j)
					break
				}
			}
		}
	}

	// top to bottom
	for i := 0; i < len(trees); i++ {
		for k := 0; k < len(trees[i]); k++ {
			for j := k + 1; j < len(trees[i]); j++ {
				if trees[j][i] >= trees[k][i] || j == len(trees[i])-1 {
					scores[k][i] = scores[k][i] * (j - k)
					break
				}
			}
		}
	}

	// bottom to top
	for i := len(trees) - 1; i >= 0; i-- {
		for k := len(trees[i]) - 1; k >= 0; k-- {
			for j := k - 1; j >= 0; j-- {
				if trees[j][i] >= trees[k][i] || j == 0 {
					scores[k][i] = scores[k][i] * (k - j)
					break
				}
			}
		}
	}

	max := 0
	for x := 0; x < len(scores); x++ {
		for y := 0; y < len(scores[x]); y++ {
			if scores[x][y] > max {
				max = scores[x][y]
			}
		}
	}
	fmt.Println(max)
}
