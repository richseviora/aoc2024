package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

const testFileName = "test.txt"
const actualFileName = "input.txt"

var LengthExceedsBlockError = fmt.Errorf("source or destination start and length exceeds length")

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

func SwapBlocks(blocks []Block, sourceStart, destStart, length int) ([]Block, error) {
	if sourceStart < 0 || destStart < 0 {
		return nil, fmt.Errorf("source or destination start less than 0")
	}
	if len(blocks) < (sourceStart+length) || len(blocks) < (destStart+length) {
		return nil, LengthExceedsBlockError
	}
	sourceEntries := blocks[sourceStart : sourceStart+length]
	destEntries := blocks[destStart : destStart+length]
	newSlice := make([]Block, len(blocks))
	copy(newSlice, blocks)
	for i, block := range sourceEntries {
		newSlice[i+destStart] = block
	}
	for i, block := range destEntries {
		newSlice[i+sourceStart] = block
	}
	return newSlice, nil
}

func CompactBlocksWholeFileId(blocks []Block, fileId int) ([]Block, bool) {
	lastFileBlockEnd := 0
	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].Id == fileId {
			lastFileBlockEnd = i
			break
		}
	}
	var lastFileBlockStart int
	for i := lastFileBlockEnd - 1; i >= 0; i-- {
		if blocks[i].IsEmpty() || blocks[i].Id != fileId {
			lastFileBlockStart = i + 1
			break
		}
	}
	requiredLength := lastFileBlockEnd - lastFileBlockStart + 1
	emptyBlockStart := -1
	emptyBlockEnd := -1
	emptyBlockFound := false
	for i := 0; i < len(blocks); i++ {
		block := blocks[i]
		if block.IsEmpty() && emptyBlockStart < 0 {
			emptyBlockStart = i
			emptyBlockEnd = i
			if requiredLength == 1 {
				emptyBlockFound = true
				break
			}
		} else if block.IsEmpty() {
			emptyBlockEnd = i
		}
		if emptyBlockStart >= 0 && emptyBlockEnd >= 0 && emptyBlockEnd-emptyBlockStart+1 >= requiredLength {
			emptyBlockFound = true
			break
		}
		if !block.IsEmpty() {
			emptyBlockStart = -1
			emptyBlockEnd = -1
			emptyBlockFound = false
		}
	}
	if emptyBlockStart >= lastFileBlockStart {
		return blocks, false
	}
	if !emptyBlockFound {
		return blocks, false
	}
	swapBlocks, err := SwapBlocks(blocks, emptyBlockStart, lastFileBlockStart, requiredLength)
	if err != nil {
		panic(err)
	}
	return swapBlocks, true
}

func CompactBlocksWholeFiles(blocks []Block) ([]Block, bool) {
	var lastFileId int
	for i := len(blocks) - 1; i >= 0; i-- {
		if !blocks[i].IsEmpty() {
			lastFileId = blocks[i].Id
			break
		}
	}
	for i := lastFileId; i >= 0; i-- {
		blocks, _ = CompactBlocksWholeFileId(blocks, i)
	}
	return blocks, false
}

func CompactUntilComplete(blocks []Block, wholeBlocks bool) []Block {
	callFunc := func(b []Block) ([]Block, bool) {
		if wholeBlocks {
			return CompactBlocksWholeFiles(blocks)
		}
		return CompactBlocks(blocks)
	}
	for {
		compacted, ok := callFunc(blocks)
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
	for _, pt2 := range []bool{false, true} {
		parsedContent := ParseFileMap(fileContent)
		compactedBlocks := CompactUntilComplete(parsedContent, pt2)
		checksum := CalculateChecksum(compactedBlocks)
		fmt.Println(fname, pt2, checksum)
	}

}

func main() {
	HandleFile(testFileName)
	HandleFile(actualFileName)
}
