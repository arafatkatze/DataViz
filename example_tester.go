package main

import (
	"fmt"

	list "github.com/Arafatk/dataviz/maps/treemap"
)

func main() {
	m := list.NewWithIntComparator()
	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")
	m.Put(1, "a") //overwrite
	fmt.Println(m.Visualizer("out.png"))
}
