package main

import (
	"fmt"

	btree "github.com/Arafatk/dataviz/maps/treebidimap"
	"github.com/emirpasic/gods/utils"
)

func main() {
	m := btree.NewWith(utils.IntComparator, utils.StringComparator)
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a")
	fmt.Println(m.Visualizer("out.png"))
}

//1 6 8 11 13 15 17 22 25 27
// stack := st.New()
// stack.Push(1)
// stack.Push(2)
// stack.Push(3)
// stack.Push("value")
// fmt.Println(stack.Visualizer("out.png"))
