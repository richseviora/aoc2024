package main

import (
	"regexp"
	"strconv"
)

type Parameter struct {
	ax, ay           int
	bx, by           int
	targetX, targetY int
}

var test = regexp.MustCompile("Button A: X\\+(\\d+), Y\\+(\\d+)\\s+Button B: X\\+(\\d+), Y\\+(\\d+)\\s+Prize: X=(\\d+), Y=(\\d+)")

func ParseInput(input string) []Parameter {
	matches := test.FindAllStringSubmatch(input, -1)
	if matches == nil {
		panic("failed to read")
	}

	parameters := make([]Parameter, len(matches))
	for i, match := range matches {
		sax, say, sbx, sby, stargetX, stargetY := match[1], match[2], match[3], match[4], match[5], match[6]
		ax, _ := strconv.Atoi(sax)
		ay, _ := strconv.Atoi(say)
		bx, _ := strconv.Atoi(sbx)
		by, _ := strconv.Atoi(sby)
		targetX, _ := strconv.Atoi(stargetX)
		targetY, _ := strconv.Atoi(stargetY)
		parameters[i] = Parameter{
			ax:      ax,
			ay:      ay,
			bx:      bx,
			by:      by,
			targetX: targetX,
			targetY: targetY,
		}
	}
	return parameters
}
