package main

import "testing"

func TestCalculation(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		ProcessChallenge("0 4 4979 24 4356119 914 85734 698829", 35)
	})
}
