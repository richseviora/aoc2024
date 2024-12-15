package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Parameter struct {
	ax, ay           int
	bx, by           int
	targetX, targetY int
	targetDelta      int
}

var test = regexp.MustCompile("Button A: X\\+(\\d+), Y\\+(\\d+)\\s+Button B: X\\+(\\d+), Y\\+(\\d+)\\s+Prize: X=(\\d+), Y=(\\d+)")

func ParseInput(input string, positionDelta int) []Parameter {
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
			ax:          ax,
			ay:          ay,
			bx:          bx,
			by:          by,
			targetX:     targetX + positionDelta,
			targetY:     targetY + positionDelta,
			targetDelta: positionDelta,
		}
	}
	return parameters
}

func (p *Parameter) Solutions() []Solution {
	aMax := p.targetX/p.ax + 1
	bMax := p.targetX/p.bx + 1
	aMin := p.targetDelta / int(math.Max(float64(p.ax), float64(p.ay)))
	bMin := p.targetDelta / int(math.Max(float64(p.bx), float64(p.by)))

	fmt.Printf("target %d %d, max %d %d, min %d %d\n", p.targetX, p.targetY, aMax, bMax, aMin, bMin)

	solutions := make([]Solution, 0)
	for a := aMin; a < aMax; a++ {
		for b := bMin; b < bMax; b++ {
			if (p.ax*a+p.bx*b == p.targetX) && (p.ay*a+p.by*b == p.targetY) {
				solutions = append(solutions, Solution{a, b})
			}
		}
	}
	return solutions
}
