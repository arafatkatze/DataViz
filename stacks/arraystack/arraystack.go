// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arraystack implements a stack backed by array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array
package arraystack

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
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

// Visualizer returns all elements in the stack (LIFO order).
func (stack *Stack) Visualizer(fileName string) (ok bool) {
	size := stack.list.Size()
	if size == 0 {
		return false
	}
	elements := make([]interface{}, size, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = stack.list.Get(i - 1) // in reverse (LIFO)
	}
	dotFileString := "digraph G {bgcolor=grey99;subgraph cluster_0 {style=filled;color=royalblue;node [style=filled,color=white, shape=rect];"
	for i := 1; i <= size; i++ {
		dotFileString = dotFileString + strconv.Itoa(i)
		dotFileString += "[fillcolor=lightpink,color=lightpink, style=filled, shape=square,label="
		switch reflect.TypeOf(elements[i-1]).Kind() {
		// Simple types
		case reflect.Bool:
			if elements[i-1].(bool) == true {
				dotFileString += "True"
			} else {
				dotFileString += "False"
			}
		case reflect.String:
			dotFileString += elements[i-1].(string)
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
			dotFileString = dotFileString + strconv.Itoa(elements[i-1].(int))
		case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
			dotFileString = dotFileString + strconv.FormatUint(elements[i-1].(uint64), 10)
		// If we've missed anything then just fmt.Sprint it
		default:
			return false
		}
		dotFileString += "];"
	}
	for i := 1; i < size; i++ {
		dotFileString += (strconv.Itoa(i) + "->")
	}
	dotFileString += (strconv.Itoa(size) + "[color=royalblue];}top[color=orange];push[color=lightpink];pop[color=lightpink];top->1[color=indianred1];1->pop[color=indianred1];push->1[color=indianred1];}")
	byteString := []byte(dotFileString)
	tmpFile, _ := ioutil.TempFile("", "TemporaryDotFile")
	tmpFile.Write(byteString)
	dotPath, _ := exec.LookPath("dot")
	dotCommandResult, _ := exec.Command(dotPath, "-Tpng", tmpFile.Name()).Output()
	ioutil.WriteFile(fileName, dotCommandResult, os.FileMode(int(0777)))
	fmt.Println(dotFileString)
	return true
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
