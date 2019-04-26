package main

import "fmt"

func main() {
	c := "adamscd"
	fmt.Println(callbackLongest(c))
	c = "addamscd"
	fmt.Println(callbackLongest(c))
}
func callbackLongest(s string) string{
	if len(s)<2{
		return  s
	}
	maxSize := 0
	startPos := 0
	for i := 0; i < len(s); i++ {
		left := i
		right := i
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxSize {
				maxSize = right - left + 1
				startPos = left
			}
			left --
			right ++
		}
		left = i
		right = i+1
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxSize {
				maxSize = right - left + 1
				startPos = left
			}
			left --
			right ++
		}
	}
	return s[startPos:maxSize]
}
