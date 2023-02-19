package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}
type Stack struct {
	top  *Node
	size int
}

func (st *Stack) IsEmpty() bool {
	return st.size == 0
}
func (st *Stack) Push(v int) {
	if st.IsEmpty() {
		node := &Node{v, nil}
		st.top = node
		return
	}
	node := &Node{v, st.top}
	st.top = node

}
func (st *Stack) pop() int {
	if st.IsEmpty() {
		return -1
	}
	node := st.top
	n := st.top.value
	st.top = node.next
	st.size--
	return n
}

func (st *Stack) Peek() int {
	if st.IsEmpty() {
		return -1
	}
	return st.top.value
}
func (st *Stack) Clear() {
	st.size = 0
	st.top = nil
}
func (st *Stack) Contains() []int {
	if st.IsEmpty() {
		return nil
	}
	var arr []int
	current := st.top
	var i = 0
	for current != nil {
		arr = append(arr, current.value)
		current = current.next
		i++
	}
	return arr
}
func (st *Stack) Increment() {
	if st.IsEmpty() {
		return
	}
	current := st.top
	for current != nil {
		current.value++
		current = current.next
	}
}
func (st *Stack) Print() {
	//n := st.size
	if st.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	current := st.top
	fmt.Print("Print: ")
	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Print("\n")

}
func (st Stack) PrintReverse() {
	if st.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	head := st.top
	var element []int
	for head != nil {
		element = append(element, head.value)
		head = head.next
	}
	fmt.Print("PrintReverse: ")
	for i := st.size - 1; i >= 0; i-- {
		fmt.Print(element[i], " ")
	}
	fmt.Print("\n")
}

func main() {
	stack := &Stack{}
	fmt.Println("Size", stack.size)
	fmt.Println("Peek", stack.Peek())
	stack.Push(5)
	stack.Push(11)
	stack.Increment()
	stack.Print()
	stack.Push(8)
	fmt.Print("Array contains: ")
	var a []int = stack.Contains()
	for i := 0; i < stack.size; i++ {
		fmt.Print(a[i], " ")
	}
	fmt.Print("\n")
	stack.Print()
	fmt.Println("Peek", stack.Peek())
	stack.PrintReverse()
	fmt.Println("Pop", stack.pop())
	stack.Print()
	stack.Clear()
	fmt.Print("After Clear: ")
	stack.Print()
	fmt.Println("Size", stack.size)
	fmt.Println("Top", stack.Peek())
}
