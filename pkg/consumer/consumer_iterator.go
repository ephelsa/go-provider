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

func (c *ConsumerIterator[T]) Remove(consumer Consumer[T]) {
	for c.HasNext() {
		if c.GetNext().ConsumerKey() == consumer.ConsumerKey() {
			if len(c.consumers) == 1 {
				c.consumers = nil
				return
			}
			c.consumers = append(c.consumers[:c.index], c.consumers[c.index+1])
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
