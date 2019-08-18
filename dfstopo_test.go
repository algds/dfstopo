package dfstopo

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	dg          DirectedGraph
	expected    []Node
	expectedErr error
}{
	{
		DirectedGraph{
			1: {
				6: struct{}{},
				7: struct{}{},
			},
			7: {
				3: struct{}{},
				9: struct{}{},
			},
			2: {
				1: struct{}{},
			},
			6: {
				7: struct{}{},
			},
			3: {
				9: struct{}{},
			},
			9: {
				5: struct{}{},
			},
		},
		[]Node{2, 1, 6, 7, 3, 9, 5},
		nil,
	},
	{
		DirectedGraph{
			1: {
				6: struct{}{},
			},
			7: {
				3: struct{}{},
				9: struct{}{},
			},
			6: {
				1: struct{}{},
			},
		},
		nil,
		ErrNotDAG(),
	},
}

func TestSort(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := Sort(tc.dg); !reflect.DeepEqual(result, tc.expected) || err != tc.expectedErr {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.expected, tc.expectedErr, result, err)
		}
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Sort(tc.dg) // nolint
		}
	}
}
