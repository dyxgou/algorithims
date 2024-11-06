package sequences

type DinamicArray[T any] struct {
	len      int
	elements []T
}

func NewDinamicArray[T any]() *DinamicArray[T] {
	return &DinamicArray[T]{
		len:      0,
		elements: make([]T, 0, 10),
	}
}

func (da *DinamicArray[T]) Len() int {
	return da.len
}

func (da *DinamicArray[T]) Append(x T) {
	da.elements = append(da.elements, x)
}

func (da *DinamicArray[T]) Iter() []T {
	return da.elements
}

func (da *DinamicArray[T]) GetAt(i int) T {
	return da.elements[i]
}

func (da *DinamicArray[T]) SetAt(i int, x T) {
	da.elements[i] = x
}

func (da *DinamicArray[T]) Reset() {
	da.len = 0
	da.elements = make([]T, 0, 10)
}
