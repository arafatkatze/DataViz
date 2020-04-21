package viz

import "errors"

// Stepper just generate the Graph at every step. The Graph list will be
// returned.
type Stepper interface {
	Next() (string, error)
	Prev() (string, error)
	Steps() ([]string, error)
	Record(string)
}

// VisualizerStepper stores all the graph steps
type VisualizerStepper struct {
	graphs  []string
	current int
}

func NewVisualizerStepper() *VisualizerStepper {
	gs := make([]string, 0)
	return &VisualizerStepper{gs, 0}
}

func (vs *VisualizerStepper) Record(g string) {
	vs.graphs = append(vs.graphs, g)
}

func (vs *VisualizerStepper) Prev() (string, error) {
	if vs.current <= 0 {
		return "", errors.New("No Previous step")
	}
	vs.current--
	return vs.graphs[vs.current], nil
}

func (vs *VisualizerStepper) Next() (string, error) {
	if vs.current >= len(vs.graphs)-1 {
		return "", errors.New("No next step")
	}
	vs.current++
	return vs.graphs[vs.current], nil
}

func (vs *VisualizerStepper) Steps() ([]string, error) {
	return vs.graphs, nil
}
