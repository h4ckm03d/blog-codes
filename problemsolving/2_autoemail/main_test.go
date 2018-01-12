package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	data := []struct {
		name string
		a    string
		b    string
		res  string
	}{
		{"Peter parkir", "Peter Parker, Peter Bajindul Parker, Peter Bajindul Parker", "Upil", "peterparker@upil.com, peterbparker@upil.com, peterbparker2@upil.com"},
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
