package sequences

type StaticSequence[T any] interface {
	Append(T)
	Len() int
	Iter() []T
	GetAt(int) T
	SetAt(int, T)
}

type DinamicSequence[T any] interface {
	StaticSequence[T]
	InsertAt(int, T) error
	DeleteAt(int, T) error
}
