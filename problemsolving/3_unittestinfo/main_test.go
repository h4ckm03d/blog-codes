package main

import "testing"

func TestPercentage(t *testing.T) {
	data := []struct {
		name    string
		test    []string
		result  []string
		percent float64
	}{
		{"50", []string{
			"test1",
			"test2a",
			"test2b",
			"test3",
		}, []string{
			"OK",
			"OK",
			"OK",
			"OK",
		},
			100},
	}

	for _, tt := range data {
		t.Run(tt.name, func(t *testing.T) {
			if Solution(tt.test, tt.result) != tt.percent {
				t.Error("salah bro")
			}
		})
	}
}
