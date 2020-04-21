package viz

import (
	"log"
	"reflect"
	"testing"

	"github.com/pennz/DataViz/trees/binaryheap"
)

func TestNewAlgVisualWrapper(t *testing.T) {
	newA := &AlgVisualWrapper{make(map[reflect.Type][]string, 0), reflect.ValueOf(nil), NewVisualizerStepper(), true}

	tests := []struct {
		name string
		want *AlgVisualWrapper
	}{
		// TODO: Add test cases.
		{"New", newA},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAlgVisualWrapper(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAlgVisualWrapper() = %v, NOT want %v", got, tt.want)
			}
		})
	}
}

func TestAlgVisualWrapper_Wrap(t *testing.T) {

	bh := binaryheap.NewWithIntComparator()

	type fields struct {
		funcs_to_wrap map[reflect.Type][]string
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
			name: "Test Wrap",
			fields: fields{
				map[reflect.Type][]string{
					reflect.TypeOf(bh): []string{"Push", "Pop"}},
				NewVisualizerStepper(),
				true},
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
			avw.Call("Pop")
			avw.Call("Pop")
			avw.Call("Push", 4)
			avw.Call("Push", 5)
			//log.Printf("%v visualize\n", avw.Call("Visualize"))
			if got != nil {
				t.Errorf("AlgVisualWrapper.Wrap() = %v, NOT want %v", got, tt.want)
			}
			log.Println(avw.Visualize())
			if avw.Visualize() == nil {
				t.Errorf("AlgVisualWrapper.Visualize() = <nil>, NOT want <nil>")
			}
		})
	}
}

func TestAlgVisualWrapper_Visualize(t *testing.T) {
	bh := binaryheap.NewWithIntComparator()

	type fields struct {
		funcs_to_wrap map[reflect.Type][]string
		stepper       *VisualizerStepper
		enabledV      bool
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Test Visualize",
			fields: fields{
				map[reflect.Type][]string{
					reflect.TypeOf(bh): []string{"Push", "Pop"}},
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
