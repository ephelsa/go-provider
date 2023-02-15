package consumer

type Iterator[T interface{}] interface {
	Append(T)
	AppendAll(...T)
	Remove(T)
	HasNext() bool
	GetNext() T
}
