# 1. Provider

This package had been inspired by [Provider by Flutter](https://pub.dev/packages/provider).

Take a look at [Provider's file](/provider/provider.go) and the official implementations below.

## 1.1. `ChangeNotifier`

[ChangeNotifier](/provider/change_notifier.go) is an implementation of `Provider` which implements:

### 1.1.1. `Watch(Consumer[T])`

Include a new `Consumer[T]` into a `Consumer[T]` iterator to be notified using `NotifyConsumers()`.

###  1.1.2. `UnWatch(Consumer[T])`

Remove the `Consumer[T]` from the `Consumer[T]` iterator.

###  1.1.3. `NotifyConsumers()`

Notify all the `Consumer[T]`s that a change has been done and the updated value is passed via `Consumer[T].Consume(T)` using the `Provider[T].Provide()` which returns a `T`.

For the example, we need to create a structure to manage the state; in this case, it will be `counter.go`.

> **`counter.go`**
> ```go
>   type Counter struct {
>       Value int16
>   }
>```

Example of `ChangeNotifier` implementation:

> **`counter_provider.go`**
>```go
> type CounterProvider interface {
> 	provider.Provider[Counter]  // Extend the interface
> 	Increment()
> 	Decrement()
> }
>```
>

> **`counter_change_notifier.go`**
>
> 1. `counterchangeNotifier` will implement `CounterProvider`. But also, we're
> extending to `provider.ChangeNotifier[Counter]`.
>
> 2. After create our `changeNotifier` reference, we need to satisfy the `ChangeNotifier.Provider` dependency too ([line 4](/provider/change_notifier.go)) with the reference that we created (remember step 1.).
>
> 3. Use `Provider.NotifyConsumers` after do a state update to notify all the consumers.
>  
> ```go
> type counterChangeNotifier struct {
> 	provider.ChangeNotifier[Counter]    // 1.
> 	Counter
> }
>
> func NewCounterChangeNotifier() CounterProvider {
> 	changeNotifier := counterChangeNotifier{
> 		provider.ChangeNotifier[Counter]{},
> 		Counter{Value: 0},
> 	}
> 
> 	changeNotifier.ChangeNotifier.Provider = &changeNotifier // 2.
> 
> 	return &changeNotifier
> }
> 
> func (p *counterChangeNotifier) Increment() {
> 	p.Counter.Value = p.Counter.Value + 1
> 	p.NotifyConsumers() // 3.
> }
> 
> func (p *counterChangeNotifier) Decrement() {
> 	p.Counter.Value = p.Counter.Value - 1
> 	p.NotifyConsumers() // 3.
> }
> 
> func (p *counterChangeNotifier) Provide() Counter {
> 	return p.Counter
> } 
> ```

# 2. Consumer

The [Consumer](/consumer/consumer.go) has to be implemented to listen any changes in a provider.

All the updates will be streamed in `Consumer[T].Consume(T)`.

An example following the last one:

> **`counter_consumer.go`**
>
> 1. In order to avoid any issues, ensure to create a different `Consumer[T].ConsumerKey` for all the implementations of `Consumer[T]` that will be listening changes of a same `Provider`.
> 
> ```go
> type counterConsumer struct{}
> 
> func NewCounterConsumer() provider.Consumer[Counter] {
> 	return &counterConsumer{}
> }
> 
> func (c *counterConsumer) Consume(counter Counter) {
> 	fmt.Printf("Counter value => %d\n", counter.Value)
> }
> 
> func (c *counterConsumer) ConsumerKey() interface{} {
> 	return struct{ Key string }{Key: "CounterConsumerKey"}  // 1
> }
> ```  
