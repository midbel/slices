package slices

import (
	"strings"
	"testing"
)

func TestRotateLeft(t *testing.T) {
	data := []struct {
		Input []string
		Want  []string
	}{
		{
			Input: []string{},
			Want:  []string{},
		},
		{
			Input: []string{"foobar"},
			Want:  []string{"foobar"},
		},
		{
			Input: []string{"foo", "bar"},
			Want:  []string{"bar", "foo"},
		},
		{
			Input: []string{"fo", "ob", "ar"},
			Want:  []string{"ob", "ar", "fo"},
		},
	}
	for i, d := range data {
		got := copyInput(d.Input)
		RotateLeft(got)

		str := strings.Join(got, ", ")
		if str != strings.Join(d.Want, ", ") {
			t.Errorf("% 2d: unexpected result! want %s, got %s", i+1, d.Want, got)
		}
	}
}

func TestRotateRight(t *testing.T) {
	data := []struct {
		Input []string
		Want  []string
	}{
		{
			Input: []string{},
			Want:  []string{},
		},
		{
			Input: []string{"foobar"},
			Want:  []string{"foobar"},
		},
		{
			Input: []string{"foo", "bar"},
			Want:  []string{"bar", "foo"},
		},
		{
			Input: []string{"fo", "ob", "ar"},
			Want:  []string{"ar", "fo", "ob"},
		},
	}
	for i, d := range data {
		got := copyInput(d.Input)
		RotateRight(got)

		str := strings.Join(got, ", ")
		if str != strings.Join(d.Want, ", ") {
			t.Errorf("% 2d: unexpected result! want %s, got %s", i+1, d.Want, got)
		}
	}
}

func copyInput(in []string) []string {
	got := make([]string, len(in))
	copy(got, in)
	return got
}
