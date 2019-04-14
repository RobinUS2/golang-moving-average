package movingaverage

import "sync"

type ConcurrentMovingAverage struct {
	ma  *MovingAverage
	mux sync.RWMutex
}

func Concurrent(ma *MovingAverage) *ConcurrentMovingAverage {
	return &ConcurrentMovingAverage{
		ma: ma,
	}
}

func (c *ConcurrentMovingAverage) Add(values ...float64) {
	c.mux.Lock()
	c.ma.Add(values...)
	c.mux.Unlock()
}

func (c *ConcurrentMovingAverage) Avg() float64 {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.ma.Avg()
}

func (c *ConcurrentMovingAverage) Min() (float64, error) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.ma.Min()
}

func (c *ConcurrentMovingAverage) Max() (float64, error) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.ma.Max()
}

func (c *ConcurrentMovingAverage) Count() int {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.ma.Count()
}

func (c *ConcurrentMovingAverage) Values() []float64 {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.ma.Values()
}

func (c *ConcurrentMovingAverage) SlotsFilled() bool {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.ma.SlotsFilled()
}
