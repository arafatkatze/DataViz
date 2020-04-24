// Package arraystack implements a stack backed by array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array
package arraystack

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/stacks"
)

func assertStackImplementation() {
	var _ stacks.Stack = (*Stack)(nil)
}

// Stack holds elements in an array-list
type Stack struct {
	list *arraylist.List
}

// New instantiates a new empty stack
func New() *Stack {
	return &Stack{list: arraylist.New()}
}

// Push adds a value onto the top of the stack
func (stack *Stack) Push(value interface{}) {
	stack.list.Add(value)
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack) Pop() (value interface{}, ok bool) {
	value, ok = stack.list.Get(stack.list.Size() - 1)
	stack.list.Remove(stack.list.Size() - 1)
	return
}

// Peek returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *Stack) Peek() (value interface{}, ok bool) {
	return stack.list.Get(stack.list.Size() - 1)
}

// Empty returns true if stack does not contain any elements.
func (stack *Stack) Empty() bool {
	return stack.list.Empty()
}

// Size returns number of elements within the stack.
func (stack *Stack) Size() int {
	return stack.list.Size()
}

// Clear removes all elements from the stack.
func (stack *Stack) Clear() {
	stack.list.Clear()
}

// Values returns all elements in the stack (LIFO order).
func (stack *Stack) Values() []interface{} {
	size := stack.list.Size()
	elements := make([]interface{}, size, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = stack.list.Get(i - 1) // in reverse (LIFO)
	}
	return elements
}

// Visualizer makes a visual image demonstrating the Stack Data Structure
// using dot language and Graphviz. It first producs a dot string corresponding
// to the Stack and then runs graphviz to output the resulting image to a file
func (stack *Stack) Visualize() string {
	size := stack.list.Size()
	if size == 0 {
		return ""
	}
	stringValues := []string{} // Putting all the elements of a stack to a string array.
	for _, value := range stack.list.Values() {
		stringValues = append(stringValues, fmt.Sprintf("%v", value))
	}
	// Adding a Dot File String that write all the contents to a dot file
	dotFileString := "digraph G {bgcolor=grey99;subgraph cluster_0 {style=filled;color=royalblue;node [style=filled,color=white, shape=rect];"
	for i := 1; i <= size; i++ {
		dotFileString += (strconv.Itoa(i) + "[fillcolor=lightpink,color=lightpink, style=filled, shape=square,label=" + stringValues[i-1] + "];")
	}
	for i := 1; i < size; i++ {
		dotFileString += (strconv.Itoa(i) + "->")
	}
	dotFileString += (strconv.Itoa(size) + "[color=royalblue];}top[color=orange];push[color=lightpink];pop[color=lightpink];top->1[color=indianred1];1->pop[color=indianred1];push->1[color=indianred1];}")
	// Writing the dot file string completed
	return dotFileString
}

// String returns a string representation of container
func (stack *Stack) String() string {
	str := "ArrayStack\n"
	values := []string{}
	for _, value := range stack.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (stack *Stack) withinRange(index int) bool {
	return index >= 0 && index < stack.list.Size()
}
