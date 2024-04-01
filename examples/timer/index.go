package timer

import (
	"fmt"
	"time"
)

type OnTick[T any] struct{}

// create ticker handler, every tick will call the func from arg then return the result to the channel
func (o *OnTick[T]) TimestampTicker(t *time.Ticker, done <-chan bool, res chan<- T, f func() T) {

	for {
		select {
		case <-done:
			close(res)
			return
		case t := <-t.C:
			go func() {
				fmt.Println("Tick at", t)
				res <- f()
			}()

		}
	}
}

func (o *OnTick[T]) StopTicker(t *time.Ticker, done chan<- bool) {
	t.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
