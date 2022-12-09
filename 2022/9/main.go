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

type Point struct {
	x int
	y int
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

	var (
		head   Point
		tails  []Point = make([]Point, 9)
		visted         = make(map[Point]struct{})
	)

	visted[Point{0, 0}] = struct{}{}

	for {
		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		split := strings.Split(string(l), " ")
		dist, _ := strconv.Atoi(split[1])
		for i := 1; i <= dist; i++ {
			switch split[0] {
			case "R":
				head = Point{x: head.x + 1, y: head.y}
			case "L":
				head = Point{x: head.x - 1, y: head.y}
			case "U":
				head = Point{x: head.x, y: head.y + 1}
			case "D":
				head = Point{x: head.x, y: head.y - 1}
			}

			var tempHead Point = head
			for i, tail := range tails {

				dx := tempHead.x - tail.x
				dy := tempHead.y - tail.y

				// if they are touching
				if (dx <= 1 && dx >= -1) && (dy <= 1 && dy >= -1) {
					//fmt.Println("do nothing")
					break
				}

				switch true {
				case dx > 0 && dy == 0:
					// move right
					tail = Point{tail.x + 1, tail.y}
				case dx > 0 && dy > 0:
					//move diagonally right up
					tail = Point{tail.x + 1, tail.y + 1}
				case dx > 0 && dy < 0:
					//move diagonally right down
					tail = Point{tail.x + 1, tail.y - 1}
				case dx < 0 && dy == 0:
					//move left
					tail = Point{tail.x - 1, tail.y}
				case dx < 0 && dy > 0:
					//move digonally left up
					tail = Point{tail.x - 1, tail.y + 1}
				case dx < 0 && dy < 0:
					//move diagonally left down
					tail = Point{tail.x - 1, tail.y - 1}
				case dx == 0 && dy > 0:
					//move up
					tail = Point{tail.x, tail.y + 1}
				case dx == 0 && dy < 0:
					//move down
					tail = Point{tail.x, tail.y - 1}
				default:
					//fmt.Println("do nothing")
				}

				tempHead = tail
				tails[i] = tail
				if i == 8 {
					visted[tail] = struct{}{}
				}

			}
			//fmt.Println(head, tails)
		}
	}
	fmt.Println(len(visted))
	//fmt.Println(visted)
}
