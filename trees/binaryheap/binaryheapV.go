// Package binaryheap implements a binary heap backed by array list.
//
// Comparator defines this heap as either min or max heap.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Binary_heap
package binaryheap

import (
	"fmt"
	"strconv"

	"github.com/Arafatk/Dataviz/trees"
	utilsRaw "github.com/Arafatk/Dataviz/utils"
	"github.com/pennz/DataViz/utils"
)

func assertTreeImplementationV() {
	var _ trees.Tree = (*HeapV)(nil)
}

// HeapV holds elements in an array-list, and for visualizer
type HeapV struct {
	*Heap
	stepper  utils.Stepper
	enabledV bool
}

// EnableV enable visualization related data
func (heap *HeapV) EnableV() {
	heap.enabledV = true
}

// SSteps why not direct use not possible?
func (heap *HeapV) SSteps() (gs []string, err error) {
	gs, err = heap.stepper.Steps()
	return
}

// NewWithV instantiates a new empty heap tree with the custom comparator.
func NewWithV(comparator utilsRaw.Comparator) *HeapV {
	return &HeapV{NewWith(comparator), utils.NewVisualizerStepper(), false}
}

// NewWithIntComparatorV instantiates a new empty heap with the IntComparator, i.e. elements are of type int.
func NewWithIntComparatorV() *HeapV {
	return &HeapV{NewWithIntComparator(), utils.NewVisualizerStepper(), false}
}

// NewWithStringComparatorV instantiates a new empty heap with the StringComparator, i.e. elements are of type string.
func NewWithStringComparatorV() *HeapV {
	return &HeapV{NewWithStringComparator(), utils.NewVisualizerStepper(), false}
}

// Push adds a value onto the heap and bubbles it up accordingly.
func (heap *HeapV) Push(values ...interface{}) {
	heap.Heap.Push(values...)
	if heap.enabledV {
		heap.stepper.Record(heap.visualize())
	}
}

// Pop removes top element on heap and returns it, or nil if heap is empty.
// Second return parameter is true, unless the heap was empty and there was nothing to pop.
func (heap *HeapV) Pop() (value interface{}, ok bool) {
	value, ok = heap.Heap.Pop()
	if heap.enabledV {
		heap.stepper.Record(heap.visualize())
	}
	return
}

// Check that the index is within bounds of the list
func (heap *HeapV) withinRange(index int) bool {
	return index >= 0 && index < heap.Heap.Size()
}

func listGet(list []interface{}, index int, size int) (interface{}, bool) {

	if !(index >= 0 && index < size) {
		return nil, false
	}
	return list[index], true
}

func (heap *HeapV) visualize() string {
	size := heap.Heap.Size()
	indexValueMap := make(map[int]interface{})
	dotString := "digraph graphname{bgcolor=white;"
	stringValues := []string{}

	list := heap.Heap.Values()

	for i := 0; i < (2 * size); i++ {
		value, exists := listGet(list, i, size)
		if exists {
			indexValueMap[i] = value // Anybody who exists is connected to parent
			if i != 0 {
				dotString += (strconv.Itoa((i-1)/2) + " -> " + strconv.Itoa((i)) + ";")
				stringValues = append(stringValues, fmt.Sprintf("%v", value))
				dotString += (strconv.Itoa(i) + "[color=steelblue1, style=filled, fillcolor = steelblue1, fontcolor=white,label=" + stringValues[len(stringValues)-1] + "];")

			} else {
				stringValues = append(stringValues, fmt.Sprintf("%v", value))
				dotString += (strconv.Itoa(i) + "[color=steelblue1, style=filled, fillcolor = steelblue1, fontcolor=white,label=" + stringValues[len(stringValues)-1] + "];")

			}
		}
	}
	dotString += "}"

	return dotString
}

// Visualizer overwrite original one by use my util, just print the string for
// debuggin
func (heap *HeapV) Visualizer(fileName string) bool {
	dotString := heap.visualize()
	return utils.WriteDotStringToPng(fileName, dotString)
}
