package main

import "fmt"

func main() {
	s := "abcaef"
	l := longofLongestSubstring(s)
	fmt.Println(l)
}

func longofLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[string]int)
	maxLength := 0
	j := 0
	for i, v := range s {
		c := string(v)
		if val, ok := m[c]; ok {
			j = Max(j, val+1)
		}
		m[c] = i
		maxLength = Max(maxLength, i-j+1)
	}
	return maxLength
}

func Max(value1, value2 int) int {
	if value1 < value2 {
		return value2
	} else {
		return value1
	}
}
