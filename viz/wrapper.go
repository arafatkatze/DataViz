package viz

import (
	"errors"
	"log"
	"reflect"

	"github.com/pennz/DataViz/trees/binaryheap"
)

type Wrapper interface {
	Wrap(interface{}) error
}

type VisualizerWrapper interface {
	Wrapper
	// Visualize after data can be collected so there is some thing to visualize
	Visualizer
}

type Visualizer interface {
	Visualize() string
}

type AlgVisualWrapper struct {
	funcs_to_wrap map[reflect.Type][]string // what needs to record
	d             reflect.Value             // wrapped datastructure
	stepper       *VisualizerStepper        // store graphs
	enabledV      bool
}

// NewAlgVisualWrapper is for generating grapsh for our datastructure
func NewAlgVisualWrapper() *AlgVisualWrapper {
	toHook := make(map[reflect.Type]([]string))
	bhp := binaryheap.NewWithIntComparator()
	toHook[reflect.TypeOf(*bhp)] = []string{"Push", "Pop"} // not possible to https://stackoverflow.com/questions/51800637/struct-type-as-map-key
	return &AlgVisualWrapper{toHook, reflect.ValueOf(nil), NewVisualizerStepper(), true}
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

func (avw *AlgVisualWrapper) Call(fname string, args ...interface{}) (out []reflect.Value) {
	//t := avw.d.Type()
	di := avw.d.Interface()

	switch t := di.(type) {
	case binaryheap.Heap: // 1. type switch , 2 different functions to hook
		dp, _ := di.(binaryheap.Heap)
		out = invoke(&dp, fname, args...)

		for _, f := range avw.funcs_to_wrap[avw.d.Type()] {
			if f == fname {
				vrv := invoke(&dp, "Visualize")[0].Interface().(string)
				avw.stepper.Record(vrv)
			}
		}
	default:
		log.Printf("Type %s not found\n", t)
	}

	return out
}

// Wrap should learn from this https://gowalker.org/reflect#MakeFunc
// So we need to creat type and its function in the runtime
// Or we need to hack to hook functions to original function in runtime
func (avw *AlgVisualWrapper) Wrap(i interface{}) error {
	_, ok := i.(Visualizer) // i is an interface wrapped a pointer to struct
	if !ok {
		return errors.New("Visualization wrap error, cannot find proper interface")
	}
	p := reflect.ValueOf(i)
	avw.d = p.Elem() // we know it is a pointer
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
	return gs
}
