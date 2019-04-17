package main

import (
	"fmt"
)
func main() {
    var cal Cal
    cal = New(5)
    cal.Do()
}

type Cal interface {
	Do()
	Left() (bool)
	Right() (bool)
	Up() (bool)
	Down() (bool)
	Move(row, col int) (bool)
}
type Number struct {
	row  int
	col  int
	n    int
	dir  int
	list [][]int
}
type Func func() bool
func (this *Number) Do() () {
	i := 0
	action := [4]Func{this.Right, this.Down, this.Left,this.Up}
	for {
		if action[this.dir%4](){
			i++
		}else {
			this.dir = this.dir+1
		}
		this.list[this.row][this.col] = i;
		if i == this.n*this.n {
			break
		}
	}
	for _,v :=range this.list{
		fmt.Println(v)
	}
}

func (this *Number) Left() (bool) {
	return this.Move(this.row, this.col-1)
}
func (this *Number) Right() (bool) {
	return this.Move(this.row, this.col+1)
}
func (this *Number) Up() (bool) {
	return this.Move(this.row-1, this.col)
}
func (this *Number) Down() (bool) {
	return this.Move(this.row+1, this.col)
}
func (this *Number) Move(row, col int) (bool) {
	if (0 <= row && row < this.n) && (0 <= col && col < this.n) {
		if this.list[row][col] == -1 {
			this.row = row
			this.col = col
			return true
		}
	}
	return false
}
func New(n int) Cal{
	list := InitAll(n)
	return &Number{
		row:  0,
		col:  -1,
		n:    n,
		dir:  0,
		list: list,
	}
}
//初始化函数
func InitAll(n int) [][]int {
	var list [][]int
	for i := 0; i < n; i++ {
		item := make([]int, 0)
		for j := 0; j < n; j++ {
			item = append(item, -1)
		}
		list = append(list, item)
	}
	return list
}
