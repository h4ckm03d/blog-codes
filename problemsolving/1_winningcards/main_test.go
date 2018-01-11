package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	data := []struct {
		name string
		a    string
		b    string
		res  int
	}{
		{"Win 3", "KQJ", "234", 3},
		{"Missmatch length", "AAT", "A", 0},
		{"RANDOM TEXT", "ADA", "ADA", 0},
	}

	for _, tc := range data {
		t.Run(tc.name, func(t *testing.T) {
			res := Solution(tc.a, tc.b)
			if res != tc.res {
				t.Errorf("Wrong result.\n==> Current result %d, must be %d", res, tc.res)
			}
		})
	}
}
