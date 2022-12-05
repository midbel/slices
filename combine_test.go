package slices

import (
	"testing"
)

func TestCombinations(t *testing.T) {
	data := []struct {
		List []string
		N    int
		Want [][]string
	}{
		{
			List: []string{"A", "B", "C", "D"},
			N:    2,
			Want: [][]string{
				{"A", "B"},
				{"A", "C"},
				{"A", "D"},
				{"B", "C"},
				{"B", "D"},
				{"C", "D"},
			},
		},
		{
			List: []string{"A", "B", "C", "D", "E"},
			N:    3,
			Want: [][]string{
				{"A", "B", "C"},
				{"A", "B", "D"},
				{"A", "B", "E"},
				{"A", "C", "D"},
				{"A", "C", "E"},
				{"A", "D", "E"},
				{"B", "C", "D"},
				{"B", "C", "E"},
				{"B", "D", "E"},
				{"C", "D", "E"},
			},
		},
	}
	for _, d := range data {
		got := Combinations(d.List, d.N)
		if !equal(d.Want, got) {
			t.Errorf("result mismatched! want %s, got %s", d.Want, got)
		}
	}
}

func TestPermutations(t *testing.T) {
	data := []struct {
		List []string
		N    int
		Want [][]string
	}{
		{
			List: []string{"A", "B", "C", "D"},
			N:    2,
			Want: [][]string{
				{"A", "B"},
				{"A", "C"},
				{"A", "D"},
				{"B", "A"},
				{"B", "C"},
				{"B", "D"},
				{"C", "A"},
				{"C", "B"},
				{"C", "D"},
				{"D", "A"},
				{"D", "B"},
				{"D", "C"},
			},
		},
		{
			List: []string{"A", "B", "C", "D"},
			N:    3,
			Want: [][]string{
				{"A", "B", "C"},
				{"A", "B", "D"},
				{"A", "C", "B"},
				{"A", "C", "D"},
				{"A", "D", "B"},
				{"A", "D", "C"},
				{"B", "A", "C"},
				{"B", "A", "D"},
				{"B", "C", "A"},
				{"B", "C", "D"},
				{"B", "D", "A"},
				{"B", "D", "C"},
				{"C", "A", "B"},
				{"C", "A", "D"},
				{"C", "B", "A"},
				{"C", "B", "D"},
				{"C", "D", "A"},
				{"C", "D", "B"},
				{"D", "A", "B"},
				{"D", "A", "C"},
				{"D", "B", "A"},
				{"D", "B", "C"},
				{"D", "C", "A"},
				{"D", "C", "B"},
			},
		},
		{
			List: []string{"A", "B", "C"},
			N:    3,
			Want: [][]string{
				{"A", "B", "C"},
				{"A", "C", "B"},
				{"B", "A", "C"},
				{"B", "C", "A"},
				{"C", "A", "B"},
				{"C", "B", "A"},
			},
		},
	}
	for _, d := range data {
		got := Permutations(d.List, d.N)
		if !equal(d.Want, got) {
			t.Errorf("result mismatched! want %s, got %s", d.Want, got)
		}
	}
}

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
