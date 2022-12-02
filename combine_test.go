package slices

import (
	"testing"
)

func TestCombine(t *testing.T) {
	data := []struct {
		Input [][]string
		Want  [][]string
	}{
		{
			Input: [][]string{{"A", "B"}},
			Want:  [][]string{{"A", "B"}},
		},
		{
			Input: [][]string{
				{"A", "B"},
				{"C"},
			},
			Want: [][]string{
				{"A", "C"},
				{"B", "C"},
			},
		},
		{
			Input: [][]string{
				{"A"},
				{"B", "C"},
			},
			Want: [][]string{
				{"A", "B"},
				{"A", "C"},
			},
		},
		{
			Input: [][]string{
				{"A", "B"},
				{"C"},
				{"D", "E", "F"},
			},
			Want: [][]string{
				{"A", "C", "D"},
				{"A", "C", "E"},
				{"A", "C", "F"},
				{"B", "C", "D"},
				{"B", "C", "E"},
				{"B", "C", "F"},
			},
		},
	}
	for _, d := range data {
		got := Combine(d.Input...)
		if !equal(d.Want, got) {
			t.Errorf("result mismatched! want %s, got %s", d.Want, got)
		}
	}
}

func equal(want, got [][]string) bool {
	if len(want) != len(got) {
		return false
	}
	for i := range want {
		if len(want[i]) != len(got[i]) {
			return false
		}
		for k := 0; k < len(want[i]); k++ {
			if want[i][k] != got[i][k] {
				return false
			}
		}
	}
	return true
}
