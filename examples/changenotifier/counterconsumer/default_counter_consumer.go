package counterconsumer

import (
	"fmt"
	"go-provider/example/changenotifier/counterchangenotifier"

	"github.com/ephelsa/go-provider/consumer"
)

type counterConsumer struct {
	tag int8
}

func NewCounterConsumer(tag int8) consumer.Consumer[counterchangenotifier.Counter] {
	return &counterConsumer{tag: tag}
}

func (c *counterConsumer) Consume(counter counterchangenotifier.Counter) {
	fmt.Printf("[%v] Counter value => %d\n", c.ConsumerKey(), counter.Value)
}

func (c *counterConsumer) ConsumerKey() interface{} {
	key := fmt.Sprintf("%d CounterConsumerKey", c.tag)

	return struct{ Key string }{Key: key}
}

func (c *counterConsumer) String() string {
	return fmt.Sprintf("Consumer %d", c.tag)
}
