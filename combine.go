package slices

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
		return [][]T{in}
	}
	for _, v := range Fst(list) {
		tmp := append(in, v)
		ret = append(ret, combine(tmp, Rest(list))...)
	}
	return ret
}
