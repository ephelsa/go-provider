package counterchangenotifier

import "github.com/ephelsa/go-provider/provider"

type counterChangeNotifier struct {
	provider.ChangeNotifier[Counter]
	Counter
}

func NewCounterChangeNotifier(initialValue int16) CounterProvider {
	changeNotifier := counterChangeNotifier{
		provider.ChangeNotifier[Counter]{},
		Counter{Value: initialValue},
	}

	changeNotifier.ChangeNotifier.Provider = &changeNotifier

	return &changeNotifier
}

func (p *counterChangeNotifier) Increment() {
	p.Counter.Value = p.Counter.Value + 1
	p.NotifyConsumers()
}

func (p *counterChangeNotifier) Decrement() {
	p.Counter.Value = p.Counter.Value - 1
	p.NotifyConsumers()
}

func (p *counterChangeNotifier) Provide() Counter {
	return p.Counter
}
