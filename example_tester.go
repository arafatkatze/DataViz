package main

import (
	"fmt"

	btree "github.com/Arafatk/dataviz/lists/doublylinkedlist"
)

func main() {
	list := btree.New()
	list.Add("a")
	list.Add("b", "c")
	list.Add("e", "h")
	//fmt.Println(tree)
	fmt.Println(list.Visualizer())
}

//1 6 8 11 13 15 17 22 25 27
// stack := st.New()
// stack.Push(1)
// stack.Push(2)
// stack.Push(3)
// stack.Push("value")
// fmt.Println(stack.Visualizer("out.png"))
