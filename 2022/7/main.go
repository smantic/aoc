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

type Node struct {
	Space    int
	Parent   *Node
	Childern map[string]*Node
}

func NewNode(p *Node) Node {
	return Node{
		Parent:   p,
		Childern: make(map[string]*Node),
	}
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
		root       = NewNode(nil)
		cur        = &root
		totalSpace int
	)

	for {

		l, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		split := strings.Split(string(l), " ")
		if split[1] == "cd" {

			if split[2] == ".." {
				cur = cur.Parent
				continue
			}

			if split[2] == "/" {
				continue
			}

			// cd <dir>
			cur = newNodeIfNotExist(cur, split[2])
			continue
		}

		i, err := strconv.Atoi(split[0])
		if len(split) == 2 && err == nil {
			cur.Space = cur.Space + i
			totalSpace = totalSpace + i
		}
	}
	fmt.Println(sum(root, 0))
	fmt.Println(result)
}

var result int = 48381165

func sum(n Node, total int) int {

	sumChildern := 0
	for _, c := range n.Childern {
		sumChildern = sumChildern + sum(*c, total)
	}

	// checky side effect to get the result.
	sumv := sumChildern + n.Space
	if sumv >= 2080344 && sumv < result {
		result = sumv
	}
	// -----------------------------

	return total + sumChildern + n.Space
}

func newNodeIfNotExist(parent *Node, name string) *Node {

	cur, exist := parent.Childern[name]
	if !exist {
		n := NewNode(parent)
		parent.Childern[name] = &n
		cur = &n
	}
	return cur
}
