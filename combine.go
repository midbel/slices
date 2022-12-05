package slices

func Combinations[T any](list []T, n int) [][]T {
	if n <= 0 || n > len(list) {
		return nil
	}
	var (
		ret  [][]T
		rest = make([]T, len(list)-1)
	)
	copy(rest, list[1:])
	for len(list) > 0 {
		ret = append(ret, combination([]T{list[0]}, rest, n-1)...)
		list = list[1:]
		if len(rest) > 0 {
			rest = rest[1:]
		}
	}
	return ret
}

func combination[T any](in []T, rest []T, n int) [][]T {
	var ret [][]T
	if n == 0 {
		return append(ret, in)
	}
	for len(rest) > 0 {
		tmp := append(in, rest[0])
		ret = append(ret, combination(tmp, rest[1:], n-1)...)
		rest = rest[1:]
	}
	return ret
}

func Permutations[T any](list []T, n int) [][]T {
	if n <= 0 || n > len(list) {
		return nil
	}
	var (
		ret  [][]T
		rest = make([]T, len(list)-1)
	)
	for i := 0; i < len(list); i++ {
		if x := copy(rest, list[:i]); x < len(rest) {
			copy(rest[x:], list[i+1:])
		}
		in := []T{list[i]}
		ret = append(ret, permute(in, rest, n-1)...)
	}
	return ret
}

func permute[T any](in []T, rest []T, n int) [][]T {
	var ret [][]T
	if n == 0 {
		return append(ret, in)
	}
	others := make([]T, len(rest)-1)
	for i := 0; i < len(rest); i++ {
		if x := copy(others, rest[:i]); x < len(others) {
			copy(others[x:], rest[i+1:])
		}
		tmp := append(in, rest[i])
		ret = append(ret, permute(tmp, others, n-1)...)
	}
	return ret
}

func Combine[T any](list ...[]T) [][]T {
	if len(list) == 0 {
		return nil
	} else if len(list) == 1 {
		return [][]T{list[0]}
	}
	var ret [][]T
	for _, v := range Fst(list) {
		ret = append(ret, combine([]T{v}, Rest(list))...)
	}
	return ret
}

func combine[T any](in []T, list [][]T) [][]T {
	var ret [][]T
	if len(list) == 0 {
		return append(ret, in)
	}
	for _, v := range Fst(list) {
		tmp := append(in, v)
		ret = append(ret, combine(tmp, Rest(list))...)
	}
	return ret
}
