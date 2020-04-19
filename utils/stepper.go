package utils

import "errors"

// Stepper just generate the Graph at every step. The Graph list will be
// returned.
type StepCollector interface {
}

type Stepper interface {
	Next() (string, error)
	Prev() (string, error)
	Steps() ([]string, error)
}

// VisualizerStepper stores all the graph steps
type VisualizerStepper struct {
	graphs  []string
	current int
}

func (vs *VisualizerStepper) Prev() (string, error) {
	if vs.current <= 0 {
		return "", errors.New("No Previous step")
	}
	vs.current -= 1
	return vs.graphs[vs.current], nil
}

func (vs *VisualizerStepper) Next() (string, error) {
	if vs.current >= len(vs.graphs)-1 {
		return "", errors.New("No next step")
	}
	vs.current += 1
	return vs.graphs[vs.current], nil
}

func (vs *VisualizerStepper) Steps() ([]string, error) {
	return vs.graphs, nil
}
