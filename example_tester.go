package main

import (
	"fmt"

	btree "github.com/Arafatk/dataviz/trees/btree"
)

func main() {
	tree := btree.NewWithIntComparator(3)
	tree.Put(1, "a")
	tree.Put(2, "b")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")

	//fmt.Println(tree)
	// BTree
	//         1
	//     2
	//         3
	// 4
	//         5
	//     6
	//         7

	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f", "g"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6, 7} (in order)

	//fmt.Println(tree)
	fmt.Println(tree.Visualizer())
	fmt.Println(tree)
}

//1 6 8 11 13 15 17 22 25 27
// stack := st.New()
// stack.Push(1)
// stack.Push(2)
// stack.Push(3)
// stack.Push("value")
// fmt.Println(stack.Visualizer("out.png"))
