package viz

type Wrapper interface {
	Wrap(interface{}) interface{}
}

type VisualizerWrapper interface {
	Wrapper
	// Visualize after data can be collected so there is some thing to visualize
	Visualizer
}

type Visualizer interface {
	Visualize() interface{}
}

type AlgVisualWrapper struct {
	funcs_to_wrap []string           // what needs to record
	stepper       *VisualizerStepper // store graphs
	enabledV      bool
}

func NewAlgVisualWrapper() *AlgVisualWrapper {
	return &AlgVisualWrapper{make([]string, 0), NewVisualizerStepper(), true}
}
func (avw *AlgVisualWrapper) Wrap(i interface{}) interface{} {
	return nil
}
func (avw *AlgVisualWrapper) Visualize() interface{} {
	return nil
}
