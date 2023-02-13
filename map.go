package slices

func MapKeys[K comparable, V any](in map[K]V) []K {
	list := make([]K, 0, len(in))
	for k := range in {
		list = append(list, k)
	}
	return list
}