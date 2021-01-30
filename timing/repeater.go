package timing

import "time"

//Repeat a function every 'interval' seconds
func Repeat(interval time.Duration, fn func()) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			fn()
			select {
			case <-time.After(interval):
			case <-stop:
				return
			}
		}
	}()

	return stop
}
