package consumer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Counter struct {
	Value int8
}

type CounterConsumer struct {
	tag int8
}

func NewCounterConsumer(tag int8) Consumer[Counter] {
	return &CounterConsumer{tag: tag}
}

func (c *CounterConsumer) Consume(counter Counter) {

}

func (c *CounterConsumer) ConsumerKey() interface{} {
	return struct{ Key string }{Key: fmt.Sprintf("TestKey#%d", c.tag)}
}

func TestConsumerIterator_Append(t *testing.T) {
	// Given
	iterator := &ConsumerIterator[Counter]{}

	c1 := NewCounterConsumer(1)
	c2 := NewCounterConsumer(2)

	// When
	iterator.Append(c1)
	iterator.Append(c2)
	iterator.Append(c1)

	// Then
	assert.Len(t, iterator.consumers, 2)
	assert.Equal(t, iterator.index, 0)
}

func TestConsumerIterator_AppendAll(t *testing.T) {
	// Given
	iterator := &ConsumerIterator[Counter]{}

	c1 := NewCounterConsumer(1)
	c2 := NewCounterConsumer(2)
	c3 := NewCounterConsumer(3)

	// When
	iterator.AppendAll(c1, c2, c3)

	// Then
	assert.Len(t, iterator.consumers, 3)
	assert.Equal(t, iterator.index, 0)
}

func TestConsumerIterator_Remove(t *testing.T) {
	// Given
	iterator := &ConsumerIterator[Counter]{}

	c1 := NewCounterConsumer(1)
	c2 := NewCounterConsumer(2)
	c3 := NewCounterConsumer(3)
	c4 := NewCounterConsumer(4)
	c5 := NewCounterConsumer(5)
	c6 := NewCounterConsumer(6)
	c7 := NewCounterConsumer(7)

	iterator.AppendAll(c1, c2, c3, c4, c5, c6)

	// When
	iterator.Remove(c3)
	iterator.Remove(c5)
	iterator.Remove(c7)

	// Then
	assert.Len(t, iterator.consumers, 4)
	assert.Contains(t, iterator.consumers, c1)
	assert.Contains(t, iterator.consumers, c2)
	assert.Contains(t, iterator.consumers, c4)
	assert.Contains(t, iterator.consumers, c6)
	assert.Equal(t, 0, iterator.index)
}

func TestConsumerIterator_HasNext(t *testing.T) {
	tests := []struct {
		name string

		index         int
		expectedIndex int

		expectedHasNext bool
	}{
		{
			name:            "When index is 0. Then HasNext should be true",
			index:           0,
			expectedIndex:   0,
			expectedHasNext: true,
		},
		{
			name:            "When index is 1. Then HasNext should be true",
			index:           1,
			expectedIndex:   1,
			expectedHasNext: true,
		},
		{
			name:            "When index is 2. Then HasNext should be false",
			index:           2,
			expectedIndex:   2,
			expectedHasNext: false,
		},
	}

	// Given
	iterator := &ConsumerIterator[Counter]{}

	c1 := NewCounterConsumer(1)
	c2 := NewCounterConsumer(2)

	iterator.AppendAll(c1, c2)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			iterator.index = tt.index

			// When
			hasNext := iterator.HasNext()

			// Then
			assert.Equal(t, tt.expectedHasNext, hasNext)
			assert.Equal(t, tt.expectedIndex, iterator.index)
		})
	}
}

func TestConsumerIterator_GetNext(t *testing.T) {
	// Given
	c1 := NewCounterConsumer(1)
	c2 := NewCounterConsumer(2)

	tests := []struct {
		name string

		index         int
		expectedIndex int

		expectedConsumer Consumer[Counter]
	}{
		{
			name:             "When index is 0. Then consumer should be c1 and index 0",
			index:            0,
			expectedIndex:    1,
			expectedConsumer: c1,
		},
		{
			name:             "When index is 1. Then consumer should be c2 and index 1",
			index:            1,
			expectedIndex:    2,
			expectedConsumer: c2,
		},
		{
			name:             "When index is 2. Then consumer should be nil and index 1",
			index:            2,
			expectedIndex:    2,
			expectedConsumer: nil,
		},
	}

	iterator := &ConsumerIterator[Counter]{}

	iterator.AppendAll(c1, c2)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			consumer := iterator.GetNext()

			// Then
			assert.Equal(t, tt.expectedConsumer, consumer)
			assert.Equal(t, tt.expectedIndex, iterator.index)
		})
	}
}
