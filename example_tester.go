package main

import (
	rbt "github.com/Arafatk/dataviz/trees/redblacktree"
)

func main() {
	tree := rbt.NewWithIntComparator()
	tree.Put(5, "e")
	tree.Put(6, "f")
	tree.Put(7, "g")
	tree.Put(3, "c")
	tree.Put(4, "d")
	tree.Put(1, "x")
	tree.Put(2, "b")
	tree.Put(1, "a") //overwrite
	tree.Visualizer("out.png")
}
