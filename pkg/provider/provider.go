package provider

import "github.com/ephelsa/provider/pkg/consumer"

type Provider[T interface{}] interface {
	Watch(consumer.Consumer[T])
	UnWatch(consumer.Consumer[T])
	NotifyConsumers()
	Provide() T
}
