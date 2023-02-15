package consumer

type Consumer[T interface{}] interface {
	Consume(T)
	ConsumerKey() interface{}
}
