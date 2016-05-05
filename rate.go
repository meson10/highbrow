package highbrow

import "time"

// Used for Rate throttling to a server
// Usage:
// limit := NewLimiter(40)
// <- limit.Start()
// limit.Stop()

func NewLimiter(rate int) *RateLimiter {
	x := RateLimiter{}
	x.SetRate(rate).SetBurst(rate)
	return &x
}

type RateLimiter struct {
	burst int //Burst Size
	rate  int //Rate per second
	tick  *time.Ticker
}

func (self *RateLimiter) SetRate(rate int) *RateLimiter {
	self.rate = rate
	return self
}

func (self *RateLimiter) SetBurst(burst int) *RateLimiter {
	self.burst = burst
	return self
}

func (self *RateLimiter) Start() <-chan time.Time {
	rate := time.Second / time.Duration(self.rate)
	self.tick = time.NewTicker(rate)

	throttle := make(chan time.Time, self.burst)

	go func() {
		for t := range self.tick.C {
			select {
			case throttle <- t:
			default:
			}
		}
	}()

	return throttle
}

func (self *RateLimiter) Stop() {
	self.tick.Stop()
}
