package main

import (
	"fmt"
	"go-provider/example/changenotifier/counterchangenotifier"
	"go-provider/example/changenotifier/counterconsumer"
)

func main() {
	changeNotifier := counterchangenotifier.NewCounterChangeNotifier(10)

	consumer1 := counterconsumer.NewCounterConsumer(1)
	consumer2 := counterconsumer.NewCounterConsumer(2)
	consumer3 := counterconsumer.NewCounterConsumer(3)

	changeNotifier.Watch(consumer1, consumer2, consumer3)

	for i := 1; i <= 10; i++ {
		if i == 5 {
			changeNotifier.Decrement()
		} else {
			changeNotifier.Increment()
		}

		if i == 3 {
			changeNotifier.UnWatch(consumer2)
		}

		if i == 6 {
			changeNotifier.Watch(consumer2)
			changeNotifier.UnWatch(consumer3)
		}

		fmt.Printf("-------------%d\n", i)
	}

	fmt.Printf("Wrapping up.\nCounter value is => %d\n", changeNotifier.Provide().Value)
}
