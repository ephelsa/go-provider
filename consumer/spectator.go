package consumer

type Spectator[T interface{}] struct {
	Key      interface{}
	Callback SpectatorFunc[T]
}

type SpectatorFunc[T interface{}] func(T)

func NewSpectator[T interface{}](key interface{}, callback SpectatorFunc[T]) Consumer[T] {
	return &Spectator[T]{
		Key:      key,
		Callback: callback,
	}
}

func (s *Spectator[T]) Consume(value T) {
	s.Callback(value)
}

func (s *Spectator[T]) ConsumerKey() interface{} {
	return s.Key
}
