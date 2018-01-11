package main

var deck = map[byte]int{
	byte('A'): 14,
	byte('K'): 13,
	byte('Q'): 12,
	byte('J'): 11,
	byte('T'): 10,
	byte('9'): 9,
	byte('8'): 8,
	byte('7'): 7,
	byte('6'): 6,
	byte('5'): 5,
	byte('4'): 4,
	byte('3'): 3,
	byte('2'): 2,
}

func Solution(A, B string) int {
	res := 0
	if len(A) != len(B) {
		return res
	}

	for idx, _ := range A {
		if value(A[idx]) > value(B[idx]) {
			res++
		}
	}

	return res
}

func value(i byte) int {
	if v, ok := deck[i]; ok {
		return v
	}

	return 0
}
