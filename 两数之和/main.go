package main

import "fmt"

func main()  {
	fmt.Println(twosum([]int{1,2,3,4,},3))
}
func twosum(nums []int,target int) []int  {
	m := make(map[int]int,0)
	for i,v:=range nums{
		if k,ok :=m[0-v];ok{
			return []int{k,i}
		}
		m[v-target] = i
	}
	return []int{}
}