package main

import (
	"fmt"
	"os"
	"slices"
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

func CompactBlocks(blocks []Block) ([]Block, bool) {
	firstEmptyBlock := slices.IndexFunc(blocks, func(b Block) bool {
		return b.IsEmpty()
	})
	if firstEmptyBlock < 0 {
		return blocks, false
	}
	var lastFileBlock int
	for i := len(blocks) - 1; i >= 0; i-- {
		if !blocks[i].IsEmpty() {
			lastFileBlock = i
			break
		}
	}
	if firstEmptyBlock >= lastFileBlock {
		return blocks, false
	}
	initialSlice := append(blocks[:firstEmptyBlock], blocks[lastFileBlock])
	return append(initialSlice, blocks[firstEmptyBlock+1:lastFileBlock]...), true
}

func CompactUntilComplete(blocks []Block) []Block {
	for {
		compacted, ok := CompactBlocks(blocks)
		if !ok {
			return compacted
		}
		blocks = compacted
	}
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

func CalculateChecksum(b []Block) int64 {
	var total int64
	for i, block := range b {
		if !block.IsEmpty() {
			total += int64(block.Id) * int64(i)
		}
	}
	return total
}

func HandleFile(fname string) {
	fileContent := ReadInput(fname)
	parsedContent := ParseFileMap(fileContent)
	compactedBlocks := CompactUntilComplete(parsedContent)
	checksum := CalculateChecksum(compactedBlocks)
	fmt.Println(fname, checksum)
}

func main() {
	HandleFile(testFileName)
	HandleFile(actualFileName)
}
