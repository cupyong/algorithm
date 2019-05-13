package main

import (
	"container/list"
	"fmt"
)

// Binary Tree
type BinaryTree struct {
	Data  interface{}
	Left  *BinaryTree
	Right *BinaryTree
}

// Constructor
func NewBinaryTree(data interface{}) *BinaryTree {
	return &BinaryTree{Data: data}
}

// 先序遍历-非递归
func (bt *BinaryTree) PreOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			res = append(res, t.Data) //visit
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}
//先序遍历 --递归
func trans(root *BinaryTree) {
	//前序遍历：先遍历根节点再遍历左子树，再遍历右子树
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	trans(root.Left)
	trans(root.Right)
}

// 中序遍历-非递归
func (bt *BinaryTree) InOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}
		if stack.Len() != 0 {
			v := stack.Back()
			t = v.Value.(*BinaryTree)
			res = append(res, t.Data) //visit
			t = t.Right
			stack.Remove(v)
		}
	}
	return res
}

func trans1(root *BinaryTree) {
	//中序遍历:先遍历左子树再遍历根，再遍历根节点再遍历右子树
	if root == nil {
		return
	}
	trans1(root.Left)
	fmt.Println(root.Data)
	trans1(root.Right)
}

// 后序遍历-非递归
func (bt *BinaryTree) PostOrderNoRecursion() []interface{} {
	t := bt
	stack := list.New()
	res := make([]interface{}, 0)
	var preVisited *BinaryTree

	for t != nil || stack.Len() != 0 {
		for t != nil {
			stack.PushBack(t)
			t = t.Left
		}

		v := stack.Back()
		top := v.Value.(*BinaryTree)

		if (top.Left == nil && top.Right == nil) || (top.Right == nil && preVisited == top.Left) || preVisited == top.Right {
			res = append(res, top.Data) //visit
			preVisited = top
			stack.Remove(v)
		} else {
			t = top.Right
		}
	}
	return res
}
func trans2(root *BinaryTree) {
	//后序遍历:先遍历左子树再遍历根，再遍历右子树，再遍历根节点
	if root == nil {
		return
	}
	trans1(root.Left)
	trans1(root.Right)
	fmt.Println(root.Data)
}

func main() {
	t := NewBinaryTree(1)
	t.Left = NewBinaryTree(3)
	t.Right = NewBinaryTree(6)
	t.Left.Left = NewBinaryTree(4)
	t.Left.Right = NewBinaryTree(5)
	t.Left.Right.Left = NewBinaryTree(8)
	t.Left.Right.Right = NewBinaryTree(9)
	t.Left.Left.Left = NewBinaryTree(7)

	fmt.Println(t.PreOrderNoRecursion())
	trans(t)
	fmt.Println(t.InOrderNoRecursion())
	trans1(t)
	fmt.Println(t.PostOrderNoRecursion())
	trans2(t)
}
