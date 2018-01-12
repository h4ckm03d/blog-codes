package main

import (
	"fmt"
	"strings"
)

func Solution(S, C string) string {

	check := map[string]int{}

	company := fmt.Sprintf("@%s.com", strings.ToLower(C))

	names := strings.Split(strings.ToLower(S), ",")

	for k, name := range names {

		n := getName(strings.TrimPrefix(name, " "))
		number := ""

		if num, ok := check[n]; !ok {
			check[n] = 1
		} else {
			check[n] = num + 1
			number = fmt.Sprintf("%d", check[n])
		}

		names[k] = fmt.Sprintf("%s%s%s", n, number, company)
	}

	return strings.Join(names, ", ")
}

func getName(s string) string {
	s = strings.Replace(s, "-", "", 0)

	splitted := strings.Split(s, " ")

	if len(s) < 2 {
		return s
	}

	lastIdx := len(splitted) - 1

	for i, v := range splitted {
		if i == 0 || i == lastIdx {
			continue
		}

		splitted[i] = string(v[0])
	}

	return strings.Join(splitted, "")
}
