package viz

import (
	"errors"
	"log"
	"reflect"
)

type Wrapper interface {
	Wrap(*interface{}) error
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
	funcs_to_wrap []string           // what needs to record
	d             *reflect.Value     // wrapped datastructure
	stepper       *VisualizerStepper // store graphs
	enabledV      bool
}

// NewAlgVisualWrapper is for generating grapsh for our datastructure
func NewAlgVisualWrapper() *AlgVisualWrapper {
	return &AlgVisualWrapper{make([]string, 0), nil, NewVisualizerStepper(), true}
}

// invoke is copied from https://stackoverflow.com/questions/8103617/call-a-struct-and-its-method-by-name-in-go
func invoke(any interface{}, name string, args ...interface{}) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	v := reflect.ValueOf(any)
	log.Println(v)
	m := v.MethodByName(name)
	log.Println(m)

	return m.Call(inputs)
}

func (avw *AlgVisualWrapper) Call(fname string, args ...interface{}) []reflect.Value {
	out := invoke(avw.d, fname, args...)
	for _, f := range avw.funcs_to_wrap {
		if f == fname {
			vrv := invoke(avw.d, "Visualize", nil)[0].Interface().(string)
			log.Println(vrv)
			avw.stepper.Record(vrv)
		}
	}
	return out
}

// Wrap should learn from this https://gowalker.org/reflect#MakeFunc
// So we need to creat type and its function in the runtime
// Or we need to hack to hook functions to original function in runtime
func (avw *AlgVisualWrapper) Wrap(i *reflect.Value) error {
	_, ok := (*i).Interface().(Visualizer)
	if !ok {
		return errors.New("Visualization wrap error, cannot find proper interface")
	}
	avw.d = i
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
