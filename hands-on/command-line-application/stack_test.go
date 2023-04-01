package main

import (
	"container/list"
	"fmt"
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	values := make([]int, 0)
	dfs(root, &values)

	return values
}

func dfs(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, values)
	*values = append(*values, root.Val)
	dfs(root.Right, values)
}

func doSomething(arr *[]int) {
	arr = append(arr, 3)
}

type Nodea struct {
	Val       int
	Neighbors []*Node
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	value := math.Min(float64(minDepth(root.Left)), float64(minDepth(root.Right)))

	return 1 + int(value)
}

func TestStack(t *testing.T) {
	var arr1 []int = []int{1, 2, 3, 4, 5}
	var arr2 = []int{1, 2, 3}
	var arr3 = []int{1}

	arr3 = append(arr3, 3)
	arr1 = append(arr1[0:3])
	arr2 = append(arr2, 100)

	fmt.Println(arr1) //output : [1, 2, 3]
	fmt.Println(arr2) //output : [1, 2, 3, 100]

	math.Min(3, 3)

	arr := make([]int, 1)
	doSomething(&arr)
	// one := Node{Val: 1}
	// two := Node{Val: 2}
	// three := Node{Val: 3}
	// four := Node{Val: 4}
	// five := Node{Val: 5}
	// seven := Node{Val: 7}

	// one.Left = &two
	// one.Right = &three
	// two.Left = &four
	// two.Right = &five
	// three.Right = &seven

	// connect(&one)

	// fmt.Print("hello")
	// temp := stack.New()

	// temp.Push(1)

	// if temp.Len() != 1 {
	// 	t.Errorf("len size: %d != 1", temp.Len())
	// }

	// value := temp.Peek()

	// if value != 1 {
	// 	t.Error("must be one")
	// }
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	doLogic(root)

	return root
}

func doLogic(root *Node) {
	queue := NewQueue()
	queue.Push(root)

	for !queue.IsEmpty() {
		size := queue.Len()

		var prev *Node
		for i := 1; i <= size; i++ {
			current := queue.Pop()

			if i != 1 {
				prev.Next = current
			}

			if current.Left != nil {
				queue.Push(current.Left)
			}

			if current.Right != nil {
				queue.Push(current.Right)
			}

			prev = current
		}
	}
}

type Queue struct {
	v *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Len() int {
	return q.v.Len()
}

func (q *Queue) IsEmpty() bool {
	if q.Len() == 0 {
		return true
	}

	return false
}

func (q *Queue) Push(v *Node) {
	q.v.PushBack(v)
}

func (q *Queue) Pop() *Node {
	front := q.v.Front()
	if front == nil {
		return nil
	}

	return q.v.Remove(front).(*Node)
}
