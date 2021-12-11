package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	// Marks are for keeping track of bingo marks.
	Marks [5][5]int
	// Values are the values on our bingo board.
	Values [5][5]int
}

func (bo *Board) String() string {
	b := strings.Builder{}
	b.WriteString("------------------------------\n")
	b.WriteString("      values \t |   marks\n")

	for i := 0; i < 5; i++ {
		s := fmt.Sprintf("%-2v | %v\n", bo.Values[i], bo.Marks[i])
		b.WriteString(s)
	}

	return b.String()
}

func main() {

	b, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("failed to open file: %v\n", err)
		return
	}

	data := string(b)

	numsidx := strings.Index(data, "\n")
	winningNums := parseWinningNums(data[:numsidx])

	boardStrs := strings.Split(
		strings.Trim(data[numsidx+1:], "\n"),
		"\n\n",
	)
	boards := make([]*Board, 0, len(boardStrs))
	for _, s := range boardStrs {
		b := parseBoard(s)
		boards = append(boards, &b)
	}

	for _, num := range winningNums {
		MarkWinningNumber(boards, num)
		winningBoards := CheckWinners(boards)
		if len(winningBoards) > 1 {
			winner := boards[winningBoards[1]]
			fmt.Printf("%s\n", winner.String())
			fmt.Printf("score: %d\n", calcScore(winner, num))
			return
		}
	}
}

func calcScore(b *Board, lastNum int) int {

	acc := 0
	for i, row := range b.Marks {
		for j, bit := range row {
			if bit == 0 {
				acc = acc + b.Values[i][j]
			}
		}
	}
	return acc * lastNum
}

func MarkWinningNumber(boards []*Board, num int) {
	for _, b := range boards {
		for i, row := range b.Values {
			for j, val := range row {
				if val == num {
					b.Marks[i][j] = 1
				}
			}
		}
	}
}

func decimalOf(in [5]int) int {

	acc := 0
	for i, bit := range in {
		acc = acc | bit<<(len(in)-1-i)
	}
	return acc
}

// Returns indexes of winning boards
func CheckWinners(boards []*Board) []int {

	winners := make([]int, 0, 10)
	for i, b := range boards {
		acc := 0
		vert1 := 0
		vert2 := 0
		for j, row := range b.Marks {
			num := decimalOf(row)
			if num == 31 {
				winners = append(winners, i)
				break
			}
			acc = acc & num
			vert1 = vert1 | row[j]
			vert2 = vert2 | row[len(row)-1-j]
		}
		if vertical(acc) {
			winners = append(winners, i)
		}
		if vert1 == 31 || vert2 == 31 {
			winners = append(winners, i)
		}
	}
	return winners
}

func vertical(acc int) bool {
	switch true {
	case acc&1 == 1,
		acc&2 == 2,
		acc&4 == 4,
		acc&8 == 8:
		return true
	}
	return false
}

func parseBoard(str string) Board {
	b := Board{}
	str = strings.Trim(str, "\n")

	rows := strings.Split(str, "\n")
	var i, j int
	for _, row := range rows {
		vals := strings.Split(row, " ")
		for _, val := range vals {

			if val == "" {
				continue
			}

			num, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
				continue
			}
			b.Values[i][j] = num
			j = j + 1
		}
		i = i + 1
		j = 0
	}
	return b
}

func parseWinningNums(str string) []int {
	splits := strings.Split(str, ",")
	nums := make([]int, 0, len(splits))
	for _, s := range splits {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("failed to parse int: %v\n", err)
			return nil
		}
		nums = append(nums, num)
	}
	return nums
}
