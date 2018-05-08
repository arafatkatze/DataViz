package main

import (
	"fmt"

	list "github.com/Arafatk/dataviz/lists/doublylinkedlist"
)

func main() {
	list := list.New()
	list.Add("a")
	list.Add("b", "c")
	fmt.Println(list.Visualizer("out.png"))
}
