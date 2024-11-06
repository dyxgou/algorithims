package sequences

type StaticSequence[T any] interface {
	build(...T) StaticSequence[T]
	len() int
	iter() []T
	get_at(int) T
	set_at(int, T)
}

type DinamicSequence[T any] interface {
	StaticSequence[T]
	insertAt(int, T) error
	deleteAt(int, T) error
}
