package main

import (
	"fmt"
	cache2 "github.com/patrickmn/go-cache"
	"slices"
	"strings"

	"github.com/kofalt/go-memoize"
)

type Sequence struct {
	Content          string
	UniqueCharacters string
}

func (s *Sequence) String() string {
	return s.Content
}

type AnalysisResult = map[Request]int

var cache = memoize.NewMemoizer(cache2.NoExpiration, cache2.NoExpiration)

func NewSequence(input string) Sequence {
	sequence := Sequence{Content: input}
	sequence.UniqueCharacters = GetUniqueValues(input)
	return sequence
}

func (s *Sequence) EqualTo(sequences []Sequence) bool {
	consolidated := ""
	for _, sequence := range sequences {
		consolidated += sequence.Content
	}
	return consolidated == s.Content
}

func SortSequences(input []Sequence) []Sequence {
	copySlice := make([]Sequence, len(input))
	copy(copySlice, input)
	cmp := func(a, b Sequence) int {
		if a.Content == b.Content {
			return 0
		} else if len(a.Content) != len(b.Content) {
			return len(a.Content) - len(b.Content)
		} else if a.Content > b.Content {
			return 1
		}
		return 0
	}
	slices.SortFunc(copySlice, cmp)
	if !slices.IsSortedFunc(copySlice, cmp) {
		panic("not sorted")
	}
	return copySlice
}

func (s *Sequence) IsSubsetOf(other Sequence) bool {
	return strings.Contains(other.Content, s.Content)
}

func CanBeComposedFrom(test string, validPermutations []Sequence) int {
	// get unique characters
	// candidates include some or all of the unique strings
	// filter for valid lengths, they cannot exceed the test length.
	fmt.Printf("CAN BE COMPOSED FROM %s %+v\n", test, validPermutations)
	filteredSequences := make(map[Sequence]bool)
	for _, sequence := range validPermutations {
		if sequence.IsSubsetOf(NewSequence(test)) {
			filteredSequences[sequence] = true
		}
	}
	if len(filteredSequences) == 0 {
		return 0
	}
	potentialPermutations := 0
	potentialMatches := make([][]Sequence, 0)
	for candidate, _ := range filteredSequences {
		//fmt.Printf("TEST: %s CANDIDATE: %+v\n", test, candidate)
		if test == candidate.Content {
			potentialPermutations++
			potentialMatches = append(potentialMatches, []Sequence{candidate})
			continue
		}
		if !strings.HasPrefix(test, candidate.Content) {
			//fmt.Printf("TEST: %s CANDIDATE: %+v FAILED\n", test, candidate)
			continue
		}
		otherResults := CanBeComposedFromMemoized(strings.TrimPrefix(test, candidate.Content), validPermutations)
		if otherResults > 0 {
			//fmt.Printf("TEST: %s CANDIDATE: %+v APPENDING: %d\n", test, candidate, otherResults)
			potentialPermutations = otherResults + potentialPermutations
			//for _, otherResult := range otherResults {
			//
			//	potentialMatches = append(potentialMatches, append([]Sequence{candidate}, otherResult...))
			//}
		}
	}

	return potentialPermutations
}

func CanBeComposedFromMemoized(test string, validPermutations []Sequence) int {
	sortedSequences := SortSequences(validPermutations)
	key := fmt.Sprintf("%s|%v", test, sortedSequences)
	result, _, cached := cache.Memoize(key, func() (interface{}, error) {
		return CanBeComposedFrom(test, sortedSequences), nil
	})
	fmt.Printf("CACHE %t %s\n", cached, key)
	return result.(int)
}

func GetUniqueValues(input string) string {
	uniqueCharacters := make(map[rune]struct{})
	for _, char := range input {
		uniqueCharacters[char] = struct{}{}
	}
	sortedRunes := make([]rune, 0, len(uniqueCharacters))
	for char := range uniqueCharacters {
		sortedRunes = append(sortedRunes, char)
	}
	slices.Sort(sortedRunes)
	stringResult := ""
	for _, char := range sortedRunes {
		stringResult += string(char)
	}
	return stringResult
}
