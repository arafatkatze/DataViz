package main

import (
	bheap "github.com/pennz/dataviz/trees/binaryheap"
)

func main() {
	heap := bheap.NewWithIntComparator()
	heap.Push(3)
	heap.Push(19)
	heap.Visualizer("heap.png")
}
