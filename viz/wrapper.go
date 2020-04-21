package viz

import (
	"log"
	"reflect"
)

type Wrapper interface {
	Wrap(interface{}) interface{}
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
	stepper       *VisualizerStepper // store graphs
	enabledV      bool
}

// NewAlgVisualWrapper is for generating grapsh for our datastructure
func NewAlgVisualWrapper() *AlgVisualWrapper {
	return &AlgVisualWrapper{make([]string, 0), NewVisualizerStepper(), true}
}

// Wrap should learn from this https://gowalker.org/reflect#MakeFunc
// So we need to creat type and its function in the runtime
// Or we need to hack to hook functions to original function in runtime
func (avw *AlgVisualWrapper) Wrap(i interface{}) interface{} {
	v, ok := i.(Visualizer)
	if !ok {
		log.Println("Visualization wrap error, cannot find proper interface")
		return nil
	}
	for _, f := range avw.funcs_to_wrap {
		vrv := v.Visualize()
		log.Println(vrv)
		//vrv, ok := vr.(string)
		if !ok {
			log.Println("Visualization wrap error, cannot find proper interface")
			return nil
		}
		avw.stepper.Record(vrv)
		reflect.TypeOf(f)
	}
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
