package provider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Consumer

type Counter struct {
	Value int8
}

type CounterConsumer struct {
	mock.Mock
	tag int8
}

func NewCounterConsumer(tag int8) *CounterConsumer {
	return &CounterConsumer{tag: tag}
}

func (c *CounterConsumer) Consume(counter Counter) {
	c.Called(counter)
}

func (c *CounterConsumer) ConsumerKey() interface{} {
	return struct{ Key string }{Key: fmt.Sprintf("TestKey#%d", c.tag)}
}

// Notifier

type CounterChangeNotifier struct {
	ChangeNotifier[Counter]
	Counter
}

type CounterProvider interface {
	Provider[Counter]
	Increment()
	Decrement()
}

func NewCounterChangeNotifier() CounterProvider {
	changeNotifier := CounterChangeNotifier{
		ChangeNotifier[Counter]{},
		Counter{Value: 10},
	}

	changeNotifier.ChangeNotifier.Provider = &changeNotifier

	return &changeNotifier
}

func NewCounterChangeNotifierWithoutProvider() CounterProvider {
	return &CounterChangeNotifier{
		ChangeNotifier[Counter]{},
		Counter{Value: 10},
	}
}

func (p *CounterChangeNotifier) Increment() {
	p.Counter.Value = p.Counter.Value + 1
	p.NotifyConsumers()
}

func (p *CounterChangeNotifier) Decrement() {
	p.Counter.Value = p.Counter.Value - 1
	p.NotifyConsumers()
}

func (p *CounterChangeNotifier) Provide() Counter {
	return p.Counter
}

// Tests
type ChangeNotifierTestSuite struct {
	suite.Suite

	CounterChangeNotifier CounterProvider
	Consumer1             *CounterConsumer
	Consumer2             *CounterConsumer
	Consumer3             *CounterConsumer
}

func (s *ChangeNotifierTestSuite) SetupTest() {
	s.CounterChangeNotifier = NewCounterChangeNotifier()

	s.Consumer1 = NewCounterConsumer(1)
	s.Consumer2 = NewCounterConsumer(2)
	s.Consumer3 = NewCounterConsumer(3)

	s.Consumer1.On("Consume", mock.AnythingOfType("Counter")).Return(nil)
	s.Consumer2.On("Consume", mock.AnythingOfType("Counter")).Return(nil)
	s.Consumer3.On("Consume", mock.AnythingOfType("Counter")).Return(nil)
}

func (s *ChangeNotifierTestSuite) TestNotifyConsumers() {
	// Given
	s.CounterChangeNotifier.Watch(s.Consumer1, s.Consumer2, s.Consumer3)

	// When
	s.CounterChangeNotifier.Increment()

	// Then
	expectedCounter := Counter{Value: 11}

	s.Consumer1.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer2.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer3.AssertCalled(s.T(), "Consume", expectedCounter)
}

func (s *ChangeNotifierTestSuite) TestUnWatchAnNotifyConsumers() {
	// Given
	s.CounterChangeNotifier.Watch(s.Consumer1, s.Consumer2, s.Consumer3)

	// When
	s.CounterChangeNotifier.UnWatch(s.Consumer2)
	s.CounterChangeNotifier.Increment()

	// Then
	expectedCounter := Counter{Value: 11}

	s.Consumer1.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer2.AssertNotCalled(s.T(), "Consume", expectedCounter)
	s.Consumer3.AssertCalled(s.T(), "Consume", expectedCounter)
}

func (s *ChangeNotifierTestSuite) TestMultipleWatchAndUnWatchAndKeepNotifyConsumers() {
	// Given
	s.CounterChangeNotifier.Watch(s.Consumer1, s.Consumer2, s.Consumer3)

	// When1
	s.CounterChangeNotifier.Increment()

	// Then1
	expectedCounter := Counter{Value: 11}

	s.Consumer1.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer2.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer3.AssertCalled(s.T(), "Consume", expectedCounter)

	// When2
	s.CounterChangeNotifier.UnWatch(s.Consumer3)
	s.CounterChangeNotifier.Increment()

	// Then2
	expectedCounter = Counter{Value: 12}

	s.Consumer1.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer2.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer3.AssertNotCalled(s.T(), "Consume", expectedCounter)

	// When3
	s.CounterChangeNotifier.UnWatch(s.Consumer1)
	s.CounterChangeNotifier.Increment()

	// Then3
	expectedCounter = Counter{Value: 13}

	s.Consumer1.AssertNotCalled(s.T(), "Consume", expectedCounter)
	s.Consumer2.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer3.AssertNotCalled(s.T(), "Consume", expectedCounter)

	// When4
	s.CounterChangeNotifier.Watch(s.Consumer1)
	s.CounterChangeNotifier.Decrement()

	// Then4
	expectedCounter = Counter{Value: 12}

	s.Consumer1.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer2.AssertCalled(s.T(), "Consume", expectedCounter)
	s.Consumer3.AssertNotCalled(s.T(), "Consume", expectedCounter)
}

func TestChangeNotifierTestSuite(t *testing.T) {
	suite.Run(t, new(ChangeNotifierTestSuite))
}

func TestNotifyConsumers_ProviderNotSet(t *testing.T) {
	// Given
	changeNotifier := NewCounterChangeNotifierWithoutProvider()
	c1 := NewCounterConsumer(1)
	changeNotifier.Watch(c1)

	// When & Then
	assert.Panics(t, changeNotifier.Increment)
}
