package main

import (
	"fmt"
	"os"
	"strconv"
)

const testFileName = "test.txt"
const actualFileName = "input.txt"

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

type Block struct {
	Id int
}

func (b Block) IsEmpty() bool {
	return b.Id < 0
}

func CreateBlocks(id int, len int) []Block {
	blocks := make([]Block, 0)
	for i := 0; i < len; i++ {
		blocks = append(blocks, Block{Id: id})
	}
	return blocks
}

func ParseFileMap(input string) []Block {
	blocks := make([]Block, 0)
	for i, digit := range input {
		length, _ := strconv.Atoi(string(digit))
		var fileId int
		if i%2 == 0 {
			fileId = i / 2
		} else {
			fileId = -1
		}
		blocks = append(blocks, CreateBlocks(fileId, length)...)

	}
	return blocks
}

func HandleFile(fname string) {
	//fileContent := ReadInput(fname)
	//parsedContent := ParseFileMap(fileContent)
}

func main() {
	HandleFile(testFileName)
	HandleFile(actualFileName)
}
