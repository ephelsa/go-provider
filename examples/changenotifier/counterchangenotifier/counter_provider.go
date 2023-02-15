package counterchangenotifier

import "github.com/ephelsa/go-provider/pkg/provider"

type CounterProvider interface {
	provider.Provider[Counter]
	Increment()
	Decrement()
}
