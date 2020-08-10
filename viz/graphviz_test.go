package viz

import (
    "reflect"
	"testing"
)

func TestDot2SVG(t *testing.T) {
    dotExample := "digraph G { A -> B } graph H { C - D }"

	tests := []struct {
		name string
        args string
		notWant string
	}{
		{"smoke", dotExample, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dot2SVG(tt.args); reflect.DeepEqual(got, tt.notWant) {
				t.Errorf("Dot2SVG() = %v, NOT want %v", got, tt.notWant)
            } else {
                t.Log(got)
            }
		})
	}
}
