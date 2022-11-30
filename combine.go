package slices

func Combine[T any](list ...[]T) [][]T {
	if len(list) == 0 {
		return nil
	} else if len(list) == 1 {
		return [][]T{list}
	}
	var ret [][]T
	for _, v := range list[0] {
		ret = append(ret, []T{v})
	}
	for _, arr := range list[1:] {
		tmp := make([][]T, len(ret)*len(arr))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = append(tmp[i], ret[i%len(ret)]...)
			tmp[i] = append(tmp[i], arr[i%len(arr)])
		}
		ret = tmp
	}
	return ret
}
