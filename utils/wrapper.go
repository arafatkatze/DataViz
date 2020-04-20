package utils

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
