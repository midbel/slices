package slices

type Ring[T any] struct {
	list []T
	size int64
	curr int64
}

func NewRing[T any](list []T) *Ring[T] {
	var (
		tmp  = make([]T, len(list))
		size = copy(tmp, list)
	)
	return &Ring[T]{
		list: tmp,
		size: int64(size),
	}
}

func (r *Ring[T]) Next() T {
	if len(r.list) == 0 {
		var zero T
		return zero
	}
	defer r.next()
	return r.list[r.curr%r.size]
}

func (r *Ring[T]) next() {
	r.curr++
	if r.curr < 0 {
		r.curr = 0
	}
}
