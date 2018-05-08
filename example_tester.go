package main

import (
	bheap "github.com/Arafatk/dataviz/trees/binaryheap"
)

func main() {
	heap := bheap.NewWithIntComparator()
	heap.Push(3)
	heap.Push(19)
	heap.Push(17)
	heap.Push(2)
	heap.Push(7)
	heap.Push(1)
	heap.Push(26)
	heap.Push(35)
	heap.Visualizer("outp.png")
	heap.Pop()
	heap.Visualizer("out.png")
	heap.Pop()
	heap.Visualizer("out1.png")
	heap.Pop()
	heap.Visualizer("out2.png")
	heap.Pop()
	heap.Visualizer("out3.png")
	heap.Pop()
	heap.Visualizer("out4.png")
	heap.Pop()
	heap.Visualizer("out5.png")
	heap.Pop()
	heap.Visualizer("out6.png")
}
