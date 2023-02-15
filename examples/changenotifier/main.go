package main

import (
	"fmt"
	"go-provider/example/changenotifier/counterchangenotifier"
	"go-provider/example/changenotifier/counterconsumer"
)

func main() {
	changeNotifier := counterchangenotifier.NewCounterChangeNotifier()

	consumer1 := counterconsumer.NewCounterConsumer(1)
	consumer2 := counterconsumer.NewCounterConsumer(2)
	consumer3 := counterconsumer.NewCounterConsumer(3)

	changeNotifier.Watch(consumer1, consumer3)

	for i := 1; i <= 10; i++ {
		if i%5 == 0 {
			changeNotifier.Decrement()
		} else {
			changeNotifier.Increment()
		}

		if i%3 == 0 {
			changeNotifier.Watch(consumer2)
		}

		if i%6 == 0 {
			changeNotifier.UnWatch(consumer3)
		}

		fmt.Println("-------------")
	}

	fmt.Printf("Wrapping up.\nCounter value is => %d\n", changeNotifier.Provide().Value)
}
