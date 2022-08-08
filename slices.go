package slices

func At[T any](args []T, ix int) T {
	var x T
	if ix < 0 || ix >= len(args) {
		return x
	}
	return args[ix]
}

func Fst[T any](args []T) T {
	return args[0]
}

func Snd[T any](args []T) T {
	return args[1]
}

func Lst[T any](args []T) T {
	return args[len(args)-1]
}

func Rest[T any](args []T) []T {
	return args[1:]
}

func Take[T any](args []T, n int) []T {
	if n >= len(args) {
		return nil
	}
	return args[n:]
}
