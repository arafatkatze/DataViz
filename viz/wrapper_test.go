package viz

import (
	"log"
	"reflect"
	"testing"

	"github.com/pennz/DataViz/lists/arraylist"
	"github.com/pennz/DataViz/lists/doublylinkedlist"
	"github.com/pennz/DataViz/lists/singlylinkedlist"
	"github.com/pennz/DataViz/maps/treemap"
	"github.com/pennz/DataViz/stacks/arraystack"
	"github.com/pennz/DataViz/trees/avltree"
	"github.com/pennz/DataViz/trees/binaryheap"
	"github.com/pennz/DataViz/trees/btree"
	"github.com/pennz/DataViz/trees/redblacktree"
)

func TestNewAlgVisualWrapper(t *testing.T) {
	newA := &AlgVisualWrapper{make(map[reflect.Type][]string, 0), reflect.ValueOf(nil), NewVisualizerStepper(), true, make(map[string]interface{}, 0)}

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

type dataH struct {
	d     interface{}
	hooks []string
}

func hookTableTest() map[reflect.Type](dataH) {

	toHook := make(map[reflect.Type](dataH))
	// not possible to https://stackoverflow.com/questions/51800637/struct-type-as-map-key
	arraylist := arraylist.New()
	doublylinkedlist := doublylinkedlist.New()
	singlylinkedlist := singlylinkedlist.New()
	// treemap := treemap.NewWith(comparator utils.Comparator)
	treemap := treemap.NewWithIntComparator()
	// treemap := treemap.NewWithStringComparator()
	arraystack := arraystack.New()
	// avltree := avltree.NewWith(comparator utils.Comparator)
	avltree := avltree.NewWithIntComparator()
	// avltree := avltree.NewWithStringComparator()
	// binaryheap := binaryheap.NewWith(comparator utils.Comparator)
	binaryheap := binaryheap.NewWithIntComparator()
	// binaryheap := binaryheap.NewWithStringComparator()
	// btree := btree.NewWith(order int, comparator utils.Comparator)
	// btree := btree.NewWithIntComparator(order int)
	btree := btree.NewWithIntComparator(4)
	// btree := btree.NewWithStringComparator(order int)
	// redblacktree := redblacktree.NewWith(comparator utils.Comparator)
	redblacktree := redblacktree.NewWithIntComparator()
	// redblacktree := redblacktree.NewWithStringComparator()
	toHook[reflect.TypeOf(arraylist)] = dataH{arraylist, []string{"Add", "Remove", "Clear", "Swap", "Insert"}}
	toHook[reflect.TypeOf(doublylinkedlist)] = dataH{doublylinkedlist, []string{"Add", "Remove", "Clear", "Swap", "Insert"}}
	toHook[reflect.TypeOf(singlylinkedlist)] = dataH{singlylinkedlist, []string{"Add", "Remove", "Clear", "Swap", "Insert"}}
	toHook[reflect.TypeOf(treemap)] = dataH{treemap, []string{"Put", "Remove", "Clear"}}
	toHook[reflect.TypeOf(arraystack)] = dataH{arraystack, []string{"Push", "Pop", "Clear"}}
	toHook[reflect.TypeOf(avltree)] = dataH{avltree, []string{"Push", "Pop", "Remove", "Clear"}}
	toHook[reflect.TypeOf(binaryheap)] = dataH{binaryheap, []string{"Push", "Pop", "Clear"}}
	toHook[reflect.TypeOf(btree)] = dataH{btree, []string{"Put", "Remove", "Clear"}}
	toHook[reflect.TypeOf(redblacktree)] = dataH{redblacktree, []string{"Put", "Remove", "Clear"}}

	return toHook
}

func TestAlgVisualWrapper_Wrap_Viz_list(t *testing.T) {
	tests := []struct {
		name string
		args *arraylist.List
		want interface{}
	}{
		{
			name: "Test Wrap and visualize",
			args: arraylist.New(),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avw := NewAlgVisualWrapper()
			got := avw.Wrap(tt.args) // already a pointer now...
			//b /Users/v/w/DataViz/viz/wrapper_test.go:60
			avw.Call("Add", 3)
			avw.Call("Add", 4)
			avw.Call("Add", 5)
			avw.Call("Swap", 0, 1)
			log.Println(avw.Visualize())
			//log.Printf("%v visualize\n", avw.Call("Visualize"))
			if got != nil {
				t.Errorf("AlgVisualWrapper.Wrap() = %v, NOT want %v", got, tt.want)
			}
			//log.Println(avw.Visualize())
			vs := avw.Visualize().([]string)
			if len(vs) != 5 {
				t.Errorf("AlgVisualWrapper.Visualize() = %v, len = %d, NOT want it", vs, len(vs))
			}
			if avw.Visualize() == nil {
				t.Errorf("AlgVisualWrapper.Visualize() = <nil>, NOT want <nil>")
			}
		})
	}
}
func TestAlgVisualWrapper_Wrap_Viz(t *testing.T) {
	tests := []struct {
		name string
		args *binaryheap.Heap
		want interface{}
	}{
		{
			name: "Test Wrap and visualize",
			args: binaryheap.NewWithIntComparator(),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			avw := NewAlgVisualWrapper()
			got := avw.Wrap(tt.args) // already a pointer now...
			//b /Users/v/w/DataViz/viz/wrapper_test.go:60
			avw.Call("Push", 3)
			avw.Call("Pop")
			avw.Call("Pop")
			avw.Call("Push", 4)
			avw.Call("Push", 5)
			//log.Println(avw.Visualize())
			//log.Printf("%v visualize\n", avw.Call("Visualize"))
			if got != nil {
				t.Errorf("AlgVisualWrapper.Wrap() = %v, NOT want %v", got, tt.want)
			}
			//log.Println(avw.Visualize())
			vs := avw.Visualize().([]string)
			if len(vs) != 5 {
				t.Errorf("AlgVisualWrapper.Visualize() = %v, len = %d, NOT want it", vs, len(vs))
			}
			if avw.Visualize() == nil {
				t.Errorf("AlgVisualWrapper.Visualize() = <nil>, NOT want <nil>")
			}
		})
	}
}

func TestAlgVisualWrapper_Wrap_Viz_batch(t *testing.T) {
	toHook := hookTableTest()

	for k, tt := range toHook {
		t.Run(k.String(), func(t *testing.T) {
			avw := NewAlgVisualWrapper()
			err := avw.Wrap(tt.d)
			if err != nil {
				t.Errorf("AlgVisualWrapper.Wrap() = %v, NOT want <ni>", err)
			} else {
				//log.Println(avw.Visualize())
				vs := avw.Visualize().([]string)
				if vs == nil {
					t.Errorf("AlgVisualWrapper.Visualize() = <nil>, NOT want <nil>")
				}
			}
		})
	}
}
