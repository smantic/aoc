package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type tree struct {
	x int
	y int
}

func main2() {

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

	var (
		line    int = 0
		visable     = make(map[tree]struct{})
		trees       = make([][]int, 99)
	)

	for i, _ := range trees {
		trees[i] = make([]int, 99)
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
	var tallest int
	for i := 0; i < len(trees); i++ {
		tallest = -1
		for j := 0; j < len(trees[i]); j++ {
			if trees[i][j] > tallest {
				tallest = trees[i][j]
				visable[tree{i, j}] = struct{}{}
			}

		}
	}

	// right to left
	for i := 0; i < len(trees); i++ {
		tallest = -1
		for j := len(trees[i]) - 1; j >= 0; j-- {
			if trees[i][j] > tallest {
				tallest = trees[i][j]
				visable[tree{i, j}] = struct{}{}
			}
		}
	}

	// top to bottom
	for i := 0; i < len(trees); i++ {
		tallest = -1
		for j := 0; j < len(trees); j++ {
			if trees[j][i] > tallest {
				tallest = trees[j][i]
				visable[tree{j, i}] = struct{}{}
			}
		}
	}

	// bottom to top
	for i := len(trees) - 1; i >= 0; i-- {
		tallest = -1
		for j := len(trees[i]) - 1; j >= 0; j-- {
			if trees[j][i] > tallest {
				tallest = trees[j][i]
				visable[tree{j, i}] = struct{}{}
			}
		}
	}
	fmt.Println(visable)
	fmt.Println(len(visable))
}
