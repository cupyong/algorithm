package main

import "fmt"

func main() {
	snake(5)
}

func snake(n int) {
	value := 1
	var intA int
	if n%2 != 0 {
		intA = n/2 + 1
	} else {
		intA = n / 2
	}
	var list [][]int
	for i := 0; i < n; i++ {
		item := make([]int, 0)
		for j := 0; j < n; j++ {
			item = append(item, -1)
		}
		list = append(list, item)
	}
	for i := 0; i < intA; i++ {
		for j := i; j < n-i; j++ {
			list[i][j] = value
			value++
		}
		for k := i + 1; k < n-i; k++ {
			list[k][n-i-1] = value
			value++
		}
		for l := n - i - 2; l >= i; l-- {
			list[n-i-1][l] = value ;
			value++
		}
		for m := n - i - 2; m > i; m-- {
			list[m][i] = value ;
			value++
		}
	}
	fmt.Println(list)
}
