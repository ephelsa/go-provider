package provider

import "github.com/ephelsa/go-provider/consumer"

type Provider[T interface{}] interface {
	Watch(...consumer.Consumer[T])
	UnWatch(consumer.Consumer[T])
	NotifyConsumers()
	Provide() T
}
