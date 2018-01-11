package main

import "testing"

func TestMax(t *testing.T) {
	// if Max(1, 3) != 3 {
	// 	t.Error("error, seharusnya 3")
	// }

	if Max(5, 3) != 5 {
		t.Error("error, seharusnya 5")
	}
}
