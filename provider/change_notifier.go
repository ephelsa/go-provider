package provider

import "github.com/ephelsa/go-provider/consumer"

type ChangeNotifier[T interface{}] struct {
	Provider[T]
	consumers consumer.ConsumerIterator[T]
}

func (c *ChangeNotifier[T]) Watch(consumers ...consumer.Consumer[T]) {
	c.consumers.AppendAll(consumers...)
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
