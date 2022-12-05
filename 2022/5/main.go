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

	var stacks [][]rune = [][]rune{
		[]rune{},
		[]rune{'H', 'T', 'Z', 'D'},
		[]rune{'Q', 'R', 'W', 'T', 'G', 'C', 'S'},
		[]rune{'P', 'B', 'F', 'Q', 'N', 'R', 'C', 'H'},
		[]rune{'L', 'C', 'N', 'F', 'H', 'Z'},
		[]rune{'G', 'L', 'F', 'Q', 'S'},
		[]rune{'V', 'P', 'W', 'Z', 'B', 'R', 'C', 'S'},
		[]rune{'Z', 'F', 'J'},
		[]rune{'D', 'L', 'V', 'Z', 'R', 'H', 'Q'},
		[]rune{'B', 'H', 'G', 'N', 'F', 'Z', 'L', 'D'},
	}

	for {

		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if len(line) == 0 || string(line)[0:4] != "move" {
			continue
		}

		split := strings.Split(string(line), " ")
		var (
			count, _ = strconv.Atoi(split[1])
			from, _  = strconv.Atoi(split[3])
			to, _    = strconv.Atoi(split[5])
			//reverse  = make([]rune, 0, 20)
		)

		fromStack := stacks[from]
		toStack := stacks[to]

		start := len(fromStack) - count
		if start <= 0 {
			start = 0
		}

		slice := fromStack[start:len(fromStack)]
		//for i := len(slice) - 1; i >= 0; i-- {
		//	reverse = append(reverse, slice[i])
		//}
		//fmt.Println(len(reverse), reverse)
		//stacks[to] = append(toStack, reverse...)

		stacks[to] = append(toStack, slice...)

		stacks[from] = fromStack[0:start]
	}

	var total int = 0
	for _, s := range stacks {
		if len(s) == 0 {
			fmt.Println(" ")
			continue
		}
		fmt.Printf("%c", s[len(s)-1])
		//fmt.Println(s)
		total += len(s)
	}

	fmt.Println()
	fmt.Println("total ", total)
}
