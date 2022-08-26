package slices

func At[T any](args []T, ix int) T {
	var x T
	if ix < 0 || ix >= len(args) {
		return x
	}
	return args[ix]
}

func Fst[T any](args []T) T {
	if len(args) == 0 {
		var x T
		return x
	}
	return args[0]
}

func Snd[T any](args []T) T {
	if len(args) < 2 {
		var x T
		return x
	}
	return args[1]
}

func Lst[T any](args []T) T {
	if len(args) == 0 {
		var x T
		return x
	}
	return args[len(args)-1]
}

func Slice[T any](args []T) []T {
	if len(args) == 0 {
		return args
	}
	return args[:len(args)-1]
}

func Rest[T any](args []T) []T {
	if len(args) == 0 {
		return nil
	}
	return args[1:]
}

func Take[T any](args []T, n int) []T {
	if n >= len(args) {
		return nil
	}
	return args[n:]
}

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

func Every[T any](args []T, fn func(T) bool) bool {
	for i := range args {
		if !fn(args[i]) {
			return false
		}
	}
	return true
}

func Some[T any](args []T, fn func(T) bool) bool {
	for i := range args {
		if fn(args[i]) {
			return true
		}
	}
	return false
}

func Reverse[T any](args []T) []T {
	list := make([]T, len(args))
	for i, j := 0, len(args)-1; i < len(args); i++ {
		list[j] = args[i]
		j--
	}
	return list
}
