package main

import (
	"math"
	"fmt"
)

func main()  {
   fmt.Println(reverse(-120))
}
func reverse(x int) int {
	var result  = 0
	for x!=0{
		result = result*10+x%10
		x = x/10
	}
	if result > math.MaxInt32 || result < -math.MaxInt32-1 {
		return 0
	}
	return  result

}


