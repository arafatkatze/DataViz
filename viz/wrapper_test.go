package viz

import (
	"reflect"
	"testing"

	"github.com/pennz/DataViz/trees/binaryheap"
)

func TestNewAlgVisualWrapper(t *testing.T) {
	newA := &AlgVisualWrapper{make([]string, 0), nil, NewVisualizerStepper(), true}
	tests := []struct {
		name string
		want *AlgVisualWrapper
	}{
		// TODO: Add test cases.
		{"New", newA},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlgVisualWrapper(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlgVisualWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlgVisualWrapper_Wrap(t *testing.T) {
	type fields struct {
		funcs_to_wrap []string
		stepper       *VisualizerStepper
		enabledV      bool
	}
	tests := []struct {
		name   string
		fields fields
		args   *binaryheap.Heap
		want   interface{}
	}{
		{
			name: "Test disableV",
			fields: fields{
				[]string{"Push", "Pop"},
				NewVisualizerStepper(),
				false},
			args: binaryheap.NewWithIntComparator(),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avw := &AlgVisualWrapper{
				funcs_to_wrap: tt.fields.funcs_to_wrap,
				stepper:       tt.fields.stepper,
				enabledV:      tt.fields.enabledV,
			}
			got := avw.Wrap(tt.args) // already a pointer now...
			//b /Users/v/w/DataViz/viz/wrapper_test.go:60
			avw.Call("Push", 3)
			avw.Call("Pop", nil)
			avw.Call("Pop", nil)
			if got != nil {
				t.Errorf("AlgVisualWrapper.Wrap() = %v, NOT want %v", got, tt.want)
			}
			if avw.Visualize() == nil {
				t.Errorf("AlgVisualWrapper.Wrap() = %v, NOT want %v", got, tt.want)
			}
		})
	}
}

func TestAlgVisualWrapper_Visualize(t *testing.T) {
	type fields struct {
		funcs_to_wrap []string
		stepper       *VisualizerStepper
		enabledV      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Test disableV",
			fields: fields{
				[]string{"Push", "Pop"},
				nil,
				false},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avw := &AlgVisualWrapper{
				funcs_to_wrap: tt.fields.funcs_to_wrap,
				stepper:       tt.fields.stepper,
				enabledV:      tt.fields.enabledV,
			}
			if got := avw.Visualize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AlgVisualWrapper.Visualize() = %v, want %v", got, tt.want)
			}
		})
	}
}
