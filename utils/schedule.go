package utils

import "time"

type Task struct {
	ticker *time.Ticker
	stop   chan bool
}

func SetInterval(fn func(), interval time.Duration, immediate bool) *Task {
	t := Task{
		time.NewTicker(interval),
		make(chan bool),
	}
	go func() {
		if immediate {
			fn()
		}
		for {
			select {
			case <-t.stop:
				return
			case <-t.ticker.C:
				fn()
			}
		}
	}()
	return &t
}

func (t Task) Stop() {
	t.stop <- true
}
