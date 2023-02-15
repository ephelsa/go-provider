package provider

import (
	"github.com/ephelsa/provider/pkg/consumer"
)

type ChangeNotifier[T interface{}] struct {
	Provider[T]
	consumers consumer.ConsumerIterator[T]
}

func (c *ChangeNotifier[T]) Watch(consumer consumer.Consumer[T]) {
	c.consumers.Append(consumer)
}

func (c *ChangeNotifier[T]) UnWatch(consumer consumer.Consumer[T]) {
	c.consumers.Remove(consumer)
}

func (c *ChangeNotifier[T]) NotifyConsumers() {
	if c.Provider == nil {
		panic("Provider intialization is required. Ensure to initialize ChangeNotifier.Provider")
	}

	iterator := c.consumers
	for iterator.HasNext() {
		iterator.GetNext().Consume(c.Provide())
	}
}
