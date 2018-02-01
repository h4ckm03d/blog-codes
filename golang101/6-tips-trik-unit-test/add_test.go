package main

import (
	"testing"
)

func TestBasic(t *testing.T) {
	if Tambah(1, 1) != 2 {
		t.Error("seharusnya 2")
	}

	if Tambah(1, 2) != 3 {
		t.Error("seharusnya 3")
	}

	// kode berulang sebanyak jumlah data test
}

func TestBasicFatal(t *testing.T) {
	if Tambah(1, 1) != 3 {
		t.Fatal("seharusnya 2")
	}

	if Tambah(1, 2) != 4 {
		t.Fatal("seharusnya 3")
	}

	// kode berulang sebanyak jumlah data test
}

func TestBasicError(t *testing.T) {
	if Tambah(1, 1) != 3 {
		t.Error("seharusnya 2")
	}

	if Tambah(1, 2) != 4 {
		t.Error("seharusnya 3")
	}

	// kode berulang sebanyak jumlah data test
}

func TestB(t *testing.T) {
	testData := []struct {
		name   string
		inputA int
		inputB int
		result int
	}{
		{"1+1", 1, 1, 2},
		{"1+2", 1, 2, 3},
		{"1+3", 1, 3, 4},
	}

	for _, tc := range testData {
		t.Run(tc.name, func(t *testing.T) {
			if Tambah(tc.inputA, tc.inputB) != tc.result {
				t.Errorf("Seharusnya %d", tc.result)
			}
		})
	}
}
