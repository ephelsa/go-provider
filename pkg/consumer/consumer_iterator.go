package consumer

type ConsumerIterator[T interface{}] struct {
	index     int
	consumers []Consumer[T]
}

func (c *ConsumerIterator[T]) Append(consumer Consumer[T]) {
	c.consumers = append(c.consumers, consumer)
}

func (c *ConsumerIterator[T]) AppendAll(consumers ...Consumer[T]) {
	for _, consumer := range consumers {
		c.Append(consumer)
	}
}

func (c *ConsumerIterator[T]) Remove(consumerToRemove Consumer[T]) {
	for i, consumer := range c.consumers {
		if consumerToRemove.ConsumerKey() == consumer.ConsumerKey() {
			tmp := make([]Consumer[T], 0)
			tmp = append(tmp, c.consumers[:i]...)
			c.consumers = append(tmp, c.consumers[i+1:]...)
		}
	}
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
