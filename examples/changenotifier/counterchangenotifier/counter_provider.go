package counterchangenotifier

import "github.com/ephelsa/go-provider/provider"

type CounterProvider interface {
	provider.Provider[Counter]
	Increment()
	Decrement()
}
