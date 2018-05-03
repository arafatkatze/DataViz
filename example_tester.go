package main

import (
	"fmt"

	st "github.com/Arafatk/dataviz/stacks/arraystack"
)

func main() {

	stack := st.New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push("value")
	fmt.Println(stack.Visualizer("out.png"))

}
