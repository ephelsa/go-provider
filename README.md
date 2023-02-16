[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Code Coverage](https://img.shields.io/badge/Code%20Coverage-100%25-success?style=flat)

# 1. Installation

```shell
go get github.com/ephelsa/go-provider
```

Check the [releases](https://github.com/ephelsa/go-provider/tags) for a specific version.

```shell
go get github.com/ephelsa/go-provider@0.0.2
```


# 2. Provider

The inspiration for this package is [Provider by Flutter](https://pub.dev/packages/provider).

Take a look at the [examples](examples/README.md).

## 2.1. `ChangeNotifier`

[ChangeNotifier](/provider/change_notifier.go) is an implementation of `Provider` which implements:

### 2.1.1. `Watch(Consumer[T])`

Include a new `Consumer[T]` into a `Consumer[T]` iterator to be notified using `NotifyConsumers()`.

###  2.1.2. `UnWatch(Consumer[T])`

Remove the `Consumer[T]` from the `Consumer[T]` iterator.

###  2.1.3. `NotifyConsumers()`

Notify all the `Consumer[T]`s that a change has been done and the updated value is passed via `Consumer[T].Consume(T)` using the `Provider[T].Provide()` which returns a `T`.

# 3. Consumer

The [Consumer](/consumer/consumer.go) has to be implemented to listen any changes in a provider.

All the updates will be streamed in `Consumer[T].Consume(T)`.

## 3.1. Spectator

`Spectator` is a `Consumer` implementation ready-to-use.

>```go
> spectator := &consumer.Spectator[T]{
>     Key: interface{},
>     Callback: func(t T) {
>         // Code
>     },
> }
>```

#

<a href="https://www.buymeacoffee.com/ephelsa"><img src="https://img.buymeacoffee.com/button-api/?text=Buy me a coffee&emoji=&slug=ephelsa&button_colour=3694ff&font_colour=ffffff&font_family=Lato&outline_colour=ffffff&coffee_colour=FFDD00" /></a>