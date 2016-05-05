package highbrow

import "sync"

func MakePool(size int, fn func(int)) *Pool {
	x := Pool{size: size, fn: fn}
	return &x
}

type Pool struct {
	size int
	fn   func(ident int)
	wg   sync.WaitGroup
}

func (self *Pool) Init() {
	for i := 0; i < self.size; i++ {
		self.wg.Add(1)
		go func(ident int) {
			self.fn(ident)
			defer self.wg.Done()
		}(i)
	}
}

func (self *Pool) Wait() {
	self.wg.Wait()
}
