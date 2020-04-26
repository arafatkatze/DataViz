package viz

import (
	"fmt"
	"log"
	"reflect"

	"github.com/pennz/DataViz/lists/arraylist"
	"github.com/pennz/DataViz/lists/doublylinkedlist"
	"github.com/pennz/DataViz/lists/singlylinkedlist"
	"github.com/pennz/DataViz/maps/treemap"
	"github.com/pennz/DataViz/stacks/arraystack"
	"github.com/pennz/DataViz/trees/avltree"
	"github.com/pennz/DataViz/trees/binaryheap"
	"github.com/pennz/DataViz/trees/btree"
	"github.com/pennz/DataViz/trees/redblacktree"
	"gonum.org/v1/gonum/graph/formats/dot"
	"gonum.org/v1/gonum/graph/formats/dot/ast"
)

type Wrapper interface {
	Wrap(interface{}) error
	Call(fname string, args ...interface{}) (out []reflect.Value)
}

type VisualizerWrapper interface {
	Wrapper
	// Visualize after data can be collected so there is some thing to visualize
	Visualizer
}

type Visualizer interface {
	Visualize() interface{}
}

type AlgVisualWrapperExtraMemory struct {
	*AlgVisualWrapper
	m interface{} // need to handle it manually.., active one and another one, and we need merge two graph
}

type AlgVisualWrapper struct {
	funcs_to_wrap  map[reflect.Type][]string // what needs to record
	d              interface{}               // wrapped datastructure
	stepper        *VisualizerStepper        // store graphs, should store call info too
	enabledV       bool
	funcCallDetail map[string]interface{} // record Get Swap detail, for visualize
}

func hookTable() map[reflect.Type]([]string) {
	toHook := make(map[reflect.Type]([]string))
	// not possible to https://stackoverflow.com/questions/51800637/struct-type-as-map-key
	arraylist := arraylist.New()
	doublylinkedlist := doublylinkedlist.New()
	singlylinkedlist := singlylinkedlist.New()
	// treemap := treemap.NewWith(comparator utils.Comparator)
	treemap := treemap.NewWithIntComparator()
	// treemap := treemap.NewWithStringComparator()
	arraystack := arraystack.New()
	// avltree := avltree.NewWith(comparator utils.Comparator)
	avltree := avltree.NewWithIntComparator()
	// avltree := avltree.NewWithStringComparator()
	// binaryheap := binaryheap.NewWith(comparator utils.Comparator)
	binaryheap := binaryheap.NewWithIntComparator()
	// binaryheap := binaryheap.NewWithStringComparator()
	// btree := btree.NewWith(order int, comparator utils.Comparator)
	// btree := btree.NewWithIntComparator(order int)
	btree := btree.NewWithIntComparator(4)
	// btree := btree.NewWithStringComparator(order int)
	// redblacktree := redblacktree.NewWith(comparator utils.Comparator)
	redblacktree := redblacktree.NewWithIntComparator()
	// redblacktree := redblacktree.NewWithStringComparator()
	toHook[reflect.TypeOf(arraylist)] = []string{"Add", "Remove", "Clear", "Swap", "Insert", "Get"} // Get to visualize sorting algs, Swap needed too
	toHook[reflect.TypeOf(doublylinkedlist)] = []string{"Add", "Remove", "Clear", "Swap", "Insert"}
	toHook[reflect.TypeOf(singlylinkedlist)] = []string{"Add", "Remove", "Clear", "Swap", "Insert"}
	toHook[reflect.TypeOf(treemap)] = []string{"Put", "Remove", "Clear"}
	toHook[reflect.TypeOf(arraystack)] = []string{"Push", "Pop", "Clear"}
	toHook[reflect.TypeOf(avltree)] = []string{"Push", "Pop", "Remove", "Clear"}
	toHook[reflect.TypeOf(binaryheap)] = []string{"Push", "Pop", "Clear"}
	toHook[reflect.TypeOf(btree)] = []string{"Put", "Remove", "Clear"}
	toHook[reflect.TypeOf(redblacktree)] = []string{"Put", "Remove", "Clear"}

	return toHook
}

// NewAlgVisualWrapper is for generating grapsh for our datastructure
func NewAlgVisualWrapper() *AlgVisualWrapper {
	toHook := hookTable()

	return &AlgVisualWrapper{toHook, reflect.ValueOf(nil), NewVisualizerStepper(), true, make(map[string]interface{}, 0)}
}

func NewAlgVisualWrapperExtraMemory() *AlgVisualWrapperExtraMemory {
	return &AlgVisualWrapperExtraMemory{NewAlgVisualWrapper(), nil}
}

// invoke is copied from https://stackoverflow.com/questions/8103617/call-a-struct-and-its-method-by-name-in-go
func invoke(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	v := reflect.ValueOf(any)
	//log.Println(v)
	m := v.MethodByName(name)
	//log.Println(m)

	return m.Call(inputs)
}

// invoke is copied from https://stackoverflow.com/questions/8103617/call-a-struct-and-its-method-by-name-in-go
func invokeWithValue(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	var m reflect.Value
	switch t := any.(type) {
	case binaryheap.Heap: // 1. type switch , 2 different functions to hook
		v := any.(binaryheap.Heap)
		m = reflect.ValueOf(&v).MethodByName(name)
	case btree.Tree: // 1. type switch , 2 different functions to hook
		v := any.(binaryheap.Heap)
		m = reflect.ValueOf(&v).MethodByName(name)
	default:
		log.Printf("Type %s not found\n", t)
	}
	return m.Call(inputs)
}

func (avw *AlgVisualWrapperExtraMemory) Call(fname string, args ...interface{}) (out []reflect.Value) {
	//t := avw.d.Type()
	//switch t := di.(type) {
	//case binaryheap.Heap: // 1. type switch , 2 different functions to hook
	//	dp, _ = &di.(binaryheap.Heap)
	//case btree.Tree: // 1. type switch , 2 different functions to hook
	//	dp, _ = di.(btree.Tree)
	//default:
	//	log.Printf("Type %s not found\n", t)
	//}
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	var m reflect.Value
	// Visualize function of different data structure

	switch t := avw.d.(type) {
	case *binaryheap.Heap: // 1. type switch , 2 different functions to hook
		v := avw.d.(*binaryheap.Heap)
		m = reflect.ValueOf(v).MethodByName(fname)
	case *btree.Tree: // 1. type switch , 2 different functions to hook
		v := avw.d.(*binaryheap.Heap)
		m = reflect.ValueOf(v).MethodByName(fname)
	case *arraylist.List: // 1. type switch , 2 different functions to hook
		v := avw.d.(*arraylist.List)
		m = reflect.ValueOf(v).MethodByName(fname)
	default:
		log.Printf("Type %s not found\n", t)
	}

	var hooked bool = false
	for _, f := range avw.funcs_to_wrap[reflect.TypeOf(avw.d)] {
		if f == fname {
			hooked = true
			break
		}
	}
	if hooked {
		avw.funcCallDetail[fname] = args
		// Call Visualize
		vrv := avw.visualize1StepBefore(fname, args...)
		if vrv != "" {
			avw.stepper.Record(vrv)
		}
		out = m.Call(inputs)
		vrv = avw.visualize1StepAfter(fname, args...)
		avw.stepper.Record(vrv)
	} else {
		out = m.Call(inputs)
	}
	return
}

// Wrap should learn from this https://gowalker.org/reflect#MakeFunc
// So we need to creat type and its function in the runtime
// Or we need to hack to hook functions to original function in runtime
func (avw *AlgVisualWrapper) Wrap(i interface{}) error {
	//_, ok := i.(Visualizer) // i is an interface wrapped a pointer to struct
	//if !ok {
	//	panic(0)
	//	//return errors.New("Visualization wrap error, cannot find proper interface")
	//}
	avw.d = i // we know it is a pointer
	return nil
}
func (avw *AlgVisualWrapper) Call(fname string, args ...interface{}) (out []reflect.Value) {
	//t := avw.d.Type()
	//switch t := di.(type) {
	//case binaryheap.Heap: // 1. type switch , 2 different functions to hook
	//	dp, _ = &di.(binaryheap.Heap)
	//case btree.Tree: // 1. type switch , 2 different functions to hook
	//	dp, _ = di.(btree.Tree)
	//default:
	//	log.Printf("Type %s not found\n", t)
	//}
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	var m reflect.Value
	// Visualize function of different data structure

	switch t := avw.d.(type) {
	case *binaryheap.Heap: // 1. type switch , 2 different functions to hook
		v := avw.d.(*binaryheap.Heap)
		m = reflect.ValueOf(v).MethodByName(fname)
	case *btree.Tree: // 1. type switch , 2 different functions to hook
		v := avw.d.(*binaryheap.Heap)
		m = reflect.ValueOf(v).MethodByName(fname)
	case *arraylist.List: // 1. type switch , 2 different functions to hook
		v := avw.d.(*arraylist.List)
		m = reflect.ValueOf(v).MethodByName(fname)
	default:
		log.Printf("Type %s not found\n", t)
	}

	var hooked bool = false
	for _, f := range avw.funcs_to_wrap[reflect.TypeOf(avw.d)] {
		if f == fname {
			hooked = true
			break
		}
	}
	if hooked {
		avw.funcCallDetail[fname] = args
		// Call Visualize
		vrv := avw.visualize1StepBefore(fname, args...)
		if vrv != "" {
			avw.stepper.Record(vrv)
		}
		out = m.Call(inputs)
		vrv = avw.visualize1StepAfter(fname, args...)
		avw.stepper.Record(vrv)
	} else {
		out = m.Call(inputs)
	}
	return
}

func (avw *AlgVisualWrapperExtraMemory) Wrap(itfc interface{}) error {
	interfaces := itfc.([]interface{})

	log.Println(interfaces)
	i := interfaces[0]
	ie := interfaces[1]
	log.Printf("%t,%v,\n %t,%v,\n %t,%v\n", interfaces, interfaces, i, i, ie, ie)

	//_, ok := i.(Visualizer) // i is an interface wrapped a pointer to struct
	//if !ok {
	//	panic(0)
	//	//return errors.New("Visualization wrap error, cannot find proper interface")
	//}
	avw.d = i // we know it is a pointer
	avw.m = ie
	return nil
}

func (avw *AlgVisualWrapper) Visualize() interface{} {
	if !avw.enabledV {
		return nil
	}

	gs, err := avw.stepper.Steps()
	if err != nil {
		log.Printf("Visualize error: %s\n", err)
		return nil
	}
	for _, g := range gs { // format for prettier print
		fmt.Println(g)
	}
	return gs
}

func (v *AlgVisualWrapperExtraMemory) visualize1StepAfter(fname string, args ...interface{}) (dotString string) {
	var nodeProp, nodeProp2 string
	var getIndex int
	var swapIdA, swapIdB int
	if fname == "Get" {
		nodeProp = "[color=black style=filled fillcolor=yellow]"
		getIndex = args[0].(int)
	}
	if fname == "Swap" {
		swapIdA, swapIdB = args[0].(int), args[1].(int)
		nodeProp = "[color=red style=filled fillcolor=red]"
		nodeProp2 = "[color=blue style=filled fillcolor=blue]"
	}

	switch t := v.d.(type) {
	case *arraylist.List:
		// Get indicate the function name
		// Swap to get us two graph, before and after swap
		values := []string{}
		dotString = "digraph graphname{bgcolor=white;subgraph cluster_0 {style=filled;color=lightgrey;node [style=filled,color=white, shape=\"Msquare\"];"
		for i, value := range v.d.(*arraylist.List).Values() {
			switch i {
			case getIndex:
				values = append(values, fmt.Sprintf("%v %s", value, nodeProp))
			case swapIdA:
				values = append(values, fmt.Sprintf("%v %s", value, nodeProp))
			case swapIdB:
				values = append(values, fmt.Sprintf("%v %s", value, nodeProp2))
			default:
				values = append(values, fmt.Sprintf("%v", value))
			}
			dotString += values[len(values)-1] + ";"
		}
		dotString += "}}"
		astFile, err := dot.ParseString(dotString)

		if err == nil {
			astFileExtra, err := dot.ParseString(visualizeList(v.m.(*arraylist.List)))
			if err == nil {

				g := astFile.Graphs[0]
				ge := astFileExtra.Graphs[0]
				clusterChangeNodeName(ge, "_", "_")
				g.Stmts = append(g.Stmts, ge.Stmts...)

				dotString = g.String()
				//fmt.Println(dotString)
			}
		}
	default:
		log.Printf("Type %s not found\n", t)
	}
	return dotString
}

func subgraphChangeNodeName(g *ast.Subgraph, pre string, ap string) {
	for _, s := range g.Stmts {
		switch s.(type) {
		case *ast.NodeStmt:
			n := s.(*ast.NodeStmt)
			n.Node.ID = pre + n.Node.ID
			log.Println(n.Node.ID)
		case *ast.Subgraph:
			sg := s.(*ast.Subgraph)
			sg.ID = sg.ID + ap
			subgraphChangeNodeName(sg, pre, ap)
		default:
		}
	}
}

func clusterChangeNodeName(g *ast.Graph, pre string, ap string) {
	for _, s := range g.Stmts {
		switch s.(type) {
		case *ast.NodeStmt:
			n := s.(*ast.NodeStmt)
			n.Node.ID = pre + n.Node.ID
			log.Println(n.Node.ID)
		case *ast.Subgraph:
			sg := s.(*ast.Subgraph)
			sg.ID = sg.ID + ap
			subgraphChangeNodeName(sg, pre, ap)
		default:
		}
	}
}

func (avw *AlgVisualWrapper) visualize1StepBefore(fname string, args ...interface{}) (dotString string) {
	var nodeProp2, nodeProp string
	var swapIdA, swapIdB int
	switch fname {
	case "Swap":
		swapIdA, swapIdB = args[0].(int), args[1].(int)
		nodeProp = "[color=red style=filled fillcolor=red]"
		nodeProp2 = "[color=blue style=filled fillcolor=blue]"
	default:
		return
	}

	switch t := avw.d.(type) {
	case *arraylist.List:
		// Get indicate the function name
		// Swap to get us two graph, before and after swap
		values := []string{}
		dotString = "digraph graphname{bgcolor=white;subgraph cluster_0 {style=filled;color=lightgrey;node [style=filled,color=white, shape=\"Msquare\"];"
		for i, value := range avw.d.(*arraylist.List).Values() {
			switch i {
			case swapIdA:
				values = append(values, fmt.Sprintf("%v %s", value, nodeProp))
			case swapIdB:
				values = append(values, fmt.Sprintf("%v %s", value, nodeProp2))
			default:
				values = append(values, fmt.Sprintf("%v", value))
			}
			dotString += values[len(values)-1] + ";"
		}
		dotString += "}}" // only one graph

		astFile, err := dot.ParseString(dotString)

		if err == nil {
			dotString = astFile.String()
			//fmt.Println(dotString)
		}
	default:
		log.Printf("Type %s not found\n", t)
	}
	return
}

func visualizeList(list *arraylist.List) string {

	values := []string{}
	dotString := "digraph graphname{bgcolor=white;subgraph cluster_0 {style=filled;color=lightgrey;node [style=filled,color=white, shape=\"Msquare\"];"
	for _, value := range list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
		dotString += values[len(values)-1] + ";"
	}
	dotString += "}}" // only one graph

	return dotString

}

// Visualizer makes a visual image demonstrating the list data structure
// using dot language and Graphviz. It first producs a dot string corresponding
// to the list and then runs graphviz to output the resulting image to a file.
func (avw *AlgVisualWrapper) visualize1StepAfter(fname string, args ...interface{}) string {
	var dotString string
	var nodeProp, nodeProp2 string
	var getIndex int
	var swapIdA, swapIdB int
	if fname == "Get" {
		nodeProp = "[color=black style=filled fillcolor=yellow]"
		getIndex = args[0].(int)
	}
	if fname == "Swap" {
		swapIdA, swapIdB = args[0].(int), args[1].(int)
		nodeProp = "[color=red style=filled fillcolor=red]"
		nodeProp2 = "[color=blue style=filled fillcolor=blue]"
	}

	switch t := avw.d.(type) {
	case *arraylist.List:
		// Get indicate the function name
		// Swap to get us two graph, before and after swap
		values := []string{}
		dotString = "digraph graphname{bgcolor=white;subgraph cluster_0 {style=filled;color=lightgrey;node [style=filled,color=white, shape=\"Msquare\"];"
		for i, value := range avw.d.(*arraylist.List).Values() {
			switch i {
			case getIndex:
				values = append(values, fmt.Sprintf("%v %s", value, nodeProp))
			case swapIdA:
				values = append(values, fmt.Sprintf("%v %s", value, nodeProp))
			case swapIdB:
				values = append(values, fmt.Sprintf("%v %s", value, nodeProp2))
			default:
				values = append(values, fmt.Sprintf("%v", value))
			}
			dotString += values[len(values)-1] + ";"
		}
		dotString += "}}"
		astFile, err := dot.ParseString(dotString)

		if err == nil {
			dotString = astFile.String()
			//fmt.Println(dotString)
		}
	default:
		log.Printf("Type %s not found\n", t)
	}
	return dotString
}
