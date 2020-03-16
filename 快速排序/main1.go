package main

import "fmt"

//查找第k大的元素
func findK(left, right int, arr []int, k int) int {
	if left >= right {
		return arr[k]
	}
	first := left
	last := right
	key := arr[first]
	for first < last {
		for first < last && arr[last] >= key {
			last -= 1
			arr[first] = arr[last]
		}
		for first < last && arr[last] <= key {
			first += 1
			arr[first] = arr[last]
		}
	}
	arr[first] = key
	if first == k {
		return arr[k]
	}else {
		if  first > k{
			return   findK(left, first, arr, k);
		}else {
			return findK(first+1, right, arr, k);
		}
	}
}
func main()  {
	list := []int{22, 56, 38, 101, 1, 18, 20, 30}
	k := findK(0, len(list) - 1,list,3)
	fmt.Println(k)
}