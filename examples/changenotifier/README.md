# 1. ChangeNotifier

Feel free to navigate to the files.

## 1.1. Struct to be managed

For this example, we are going to create an struct which will contain a `Value` to hold and notify any changes.

> **[`counter.go`](./counterchangenotifier/counter.go)**
> ```go
>   type Counter struct {
>       Value int16
>   }
>```

## 1.2. `provider.Provider[T]` implementation

In this case, `counter_provider.go` is the interface which implement `provider.Provider[Counter]` and contain some other functions that will modify the state of a `Counter`.

> **[`counter_provider.go`](./counterchangenotifier/counter_provider.go)**
>```go
> type CounterProvider interface {
> 	provider.Provider[Counter]  // Extend the interface
> 	Increment()
> 	Decrement()
> }
>```
>

## 1.3. `provider.ChangeNotifier[T]` implementation

Then, we have to implement `CounterProvider` which implements `provider.Provider[Counter]` extending to `provider.ChangeNotifier[Counter]`.

> **[`counter_change_notifier.go`](./counterchangenotifier/counter_change_notifier.go)**
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

## 1.4. `Consumer` implementation

To observe the changes of `CounterChangeNotifier` we have to implement at least one `Consumer`.

> **[`default_counter_consumer.go`](./counterconsumer/default_counter_consumer.go)**
>
> 1. In order to avoid any issues, ensure to create a different `Consumer[T].ConsumerKey` for all the implementations of `Consumer[T]` that will be listening changes of a same `Provider`.
> 
> ```go
> type counterConsumer struct{}
> 
> func NewCounterConsumer() consumer.Consumer[counterchangenotifier.Counter] {
> 	return &counterConsumer{}
> }
> 
> func (c *counterConsumer) Consume(counter counterchangenotifier.Counter) {
> 	fmt.Printf("[%v] Counter value => %d\n", c.ConsumerKey(), counter.Value)
> }
> 
> func (c *counterConsumer) ConsumerKey() interface{} {
> 	return struct{ Key string }{Key: "CounterConsumerKey"}  // 1
> }
> ```  
