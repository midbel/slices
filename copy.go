package slices

func CopyMap[T comparable, V any](arr map[T]V) map[T]V {
	ret := make(map[T]V)
	for k, v := range arr {
		ret[k] = v
	}
	return ret
}

func AppendValues[T any](arr [][]T, values []T) [][]T {
	if len(arr) == 0 {
		for i := range values {
			arr = append(arr, []T{values[i]})
		}
		return arr
	}
	var (
		old = copyValues(arr)
		ret [][]T
	)
	for i := range values {
		arr := copyValues(old)
		for j := range arr {
			arr[j] = append(arr[j], values[i])
		}
		ret = append(ret, arr...)
	}
	return ret
}

func copyValues[T any](arr [][]T) [][]T {
	var ret [][]T
	for i := range arr {
		tmp := make([]T, len(arr[i]))
		copy(tmp, arr[i])
		ret = append(ret, tmp)
	}
	return ret
}