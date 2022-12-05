package slices

import (
	"math/rand"
)

// At returns the element at the index of the given slice. At returns the zero
// value of the slice type if index is out of range
func At[T any](args []T, ix int) T {
	var x T
	if ix < 0 || ix >= len(args) {
		return x
	}
	return args[ix]
}

// Fst returns the first element of the given slice. It returns the zero value
// of the slice type if the slice is empty
func Fst[T any](args []T) T {
	if len(args) == 0 {
		var x T
		return x
	}
	return args[0]
}

// Snd returns the second element of the given slice. It returns the zero value
// of the slice type if the length of the given slice is less than two elements
func Snd[T any](args []T) T {
	if len(args) < 2 {
		var x T
		return x
	}
	return args[1]
}

// Lst returns the last element of the given slice. It returns the zero value of
// the slice type if the given slice is empty
func Lst[T any](args []T) T {
	if len(args) == 0 {
		var x T
		return x
	}
	return args[len(args)-1]
}

// Slice returns the given slice with the last element removed from it
func Slice[T any](args []T) []T {
	if len(args) == 0 {
		return args
	}
	return args[:len(args)-1]
}

// Rest returns the given slice with the first element removed from it
func Rest[T any](args []T) []T {
	return Take(args, 1)
}

// Take returns the given slice with the n first element removed from it. If
// n is greater than the size of the given slice, an empty slice is returned
func Take[T any](args []T, n int) []T {
	if n < 0 || n >= len(args) {
		return nil
	}
	return args[n:]
}

func Splice[T any](args []T, i, n int, rest ...T) []T {
	if i < 0 || i >= len(args) {
		return nil
	}
	var list []T
	list = append(list, args[:i]...)
	if len(rest) > 0 {
		list = append(list, rest...)
	}
	if i+n < len(args) {
		list = append(list, args[i+n:]...)
	}
	return list
}

// Index returns the index of the first element for which the given function
// returns true
func Index[T any](args []T, fn func(T) bool) int {
	for i := range args {
		if fn(args[i]) {
			return i
		}
	}
	return -1
}

// Filter returns a new slice with all elements from the input slice for which
// the given function returns true
func Filter[T any](args []T, fn func(T) bool) []T {
	list := make([]T, 0, len(args))
	for i := range args {
		if !fn(args[i]) {
			continue
		}
		list = append(list, args[i])
	}
	return list
}

// Every returns true if all the elements of the input slice pass the test
// of the given function. Every stops iterating as soon as one of the element
// given to the function fails
func Every[T any](args []T, fn func(T) bool) bool {
	for i := range args {
		if !fn(args[i]) {
			return false
		}
	}
	return true
}

// Some returns true if at least one element of the input slice pass the test
// of the given function. Every stops iterating as soon as one of the element
// given to the function succeed
func Some[T any](args []T, fn func(T) bool) bool {
	for i := range args {
		if fn(args[i]) {
			return true
		}
	}
	return false
}

// Reverse returns a new slice with all the elements of the input slice in
// reverse order
func Reverse[T any](args []T) []T {
	list := make([]T, len(args))
	for i, j := 0, len(args)-1; i < len(args); i++ {
		list[j] = args[i]
		j--
	}
	return list
}

// Shuffle randomizes the elements of the slice in place
func Shuffle[T any](args []T) []T {
	rand.Shuffle(len(args), func(i, j int) {
		args[i], args[j] = args[j], args[i]
	})
	return args
}

// Foreach gives each elements of the given slices to a function
func Foreach[T any](args []T, do func(T)) {
	for i := range args {
		do(args[i])
	}
}

// Apply call the given function to each element of the slice and set the
// returned value at the index of the input element
func Apply[T any](args []T, do func(T) T) {
	for i := range args {
		args[i] = do(args[i])
	}
}

// Map call the given function to each element of the slice and set the
// returned value at the index of the input element in a copy of the input
// slice
func Map[T any](args []T, do func(T) T) []T {
	vs := make([]T, len(args))
	copy(vs, args)
	Apply(vs, do)
	return vs
}

// Prepend appends an element at the beginning of the input slice
func Prepend[T any](fst T, rest []T) []T {
	arr := []T{fst}
	return append(arr, rest...)
}

func RotateLeft[T any](list []T) {
	if len(list) <= 1 {
		return
	}
	var (
		lst = Lst(list)
		fst = Fst(list)
	)
	for i := len(list) - 1; i > 0; i-- {
		list[i-1], lst = lst, list[i-1]
	}
	list[len(list)-1] = fst
}

func RotateRight[T any](list []T) {
	if len(list) <= 1 {
		return
	}
	var (
		lst = Lst(list)
		fst = Fst(list)
	)
	for i := 0; i < len(list)-1; i++ {
		list[i+1], fst = fst, list[i+1]
	}
	list[0] = lst
}
