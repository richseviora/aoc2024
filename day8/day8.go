package day8

import (
	"fmt"
	"os"
	"strings"
)

type Grid struct {
	Nodes []Node
}

type Node struct {
	Frequency string
	Row       int
	Column    int
}

func NewTableFromString(input string) Grid {
	rows := strings.Split(input, "\n")
	for ri, row := range rows {
		for ci, column := range strings.Split(row, "") {

		}
	}
}

const filename = "input.txt"

func ReadInput(fname string) string {
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		fmt.Printf("error received: %e", err)
		panic(err)
	}

	content, err := os.ReadFile(fname)
	if err != nil {
		fmt.Printf("error received: %v", err)
		panic(err)
	}
	return string(content)
}

func main() {

}
