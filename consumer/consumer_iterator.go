package consumer

type ConsumerIterator[T interface{}] struct {
	index     int
	consumers []Consumer[T]
}

func (c *ConsumerIterator[T]) Append(consumer Consumer[T]) {
	if exist, _ := c.existInConsumers(consumer); !exist {
		c.consumers = append(c.consumers, consumer)
	}
}

func (c *ConsumerIterator[T]) AppendAll(consumers ...Consumer[T]) {
	for _, consumer := range consumers {
		c.Append(consumer)
	}
}

func (c *ConsumerIterator[T]) Remove(consumerToRemove Consumer[T]) {
	exist, index := c.existInConsumers(consumerToRemove)
	if exist {
		tmp := make([]Consumer[T], 0)
		tmp = append(tmp, c.consumers[:index]...)
		c.consumers = append(tmp, c.consumers[index+1:]...)
	}
}

func (c *ConsumerIterator[T]) existInConsumers(consumerToFind Consumer[T]) (bool, int) {
	for index, consumer := range c.consumers {
		if exist := consumer.ConsumerKey() == consumerToFind.ConsumerKey(); exist {
			return exist, index
		}
	}

	return false, -1
}

func (c *ConsumerIterator[T]) HasNext() bool {
	return c.index < len(c.consumers)
}

func (c *ConsumerIterator[T]) GetNext() Consumer[T] {
	if c.HasNext() {
		consumer := c.consumers[c.index]
		c.index++
		return consumer
	}

	return nil
}
