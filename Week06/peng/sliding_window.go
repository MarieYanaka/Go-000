package peng
import (
	"sync"
	"time"
)

// SlidingWindow contains lots of windows.
type SlidingWindow struct {
	Windows map[int64]*window
	Mutex   *sync.RWMutex
}

// window stores value.
type window struct {
	Value float64
}

// NewSlidingWindow return a new sliding window.
func NewSlidingWindow() *SlidingWindow {
	return &SlidingWindow{
		Windows: make(map[int64]*window),
		Mutex:   &sync.RWMutex{},
	}
}

// getCurrentWindow return current window of time.Now().
func (sw *SlidingWindow) getCurrentWindow() *window {
	now := time.Now().Unix()

	w, ok := sw.Windows[now]
	if !ok {
		w = &window{}
		sw.Windows[now] = w
	}
	return w
}

// removeOldWindows remove 10 second before windows.
func (sw *SlidingWindow) removeOldWindows() {
	now := time.Now().Unix() - 10

	for timestamp := range sw.Windows {
		if timestamp <= now {
			delete(sw.Windows, timestamp)
		}
	}
}

// Increment add i to current window.
func (sw *SlidingWindow) Increment(i float64) {
	if i == 0 {
		return
	}
	sw.Mutex.RLock()
	defer sw.Mutex.RUnlock()

	w := sw.getCurrentWindow()
	w.Value += i
	sw.removeOldWindows()
}

// UpdateMax updates the maximum value in the current window.
func (sw *SlidingWindow) UpdateMax(n float64) {
	sw.Mutex.Lock()
	defer sw.Mutex.Unlock()

	w := sw.getCurrentWindow()
	if n > w.Value {
		w.Value = n
	}
	sw.removeOldWindows()
}

// Sum sums the values over the windows in the last 10 seconds.
func (sw *SlidingWindow) Sum(now time.Time) float64 {
	var sum float64

	sw.Mutex.RLock()
	defer sw.Mutex.RUnlock()

	for timestamp, window := range sw.Windows {
		if timestamp >= now.Unix()-10 {
			sum += window.Value
		}
	}

	return sum
}

// Max returns the maximum value seen in the last 10 seconds.
func (sw *SlidingWindow) Max(now time.Time) float64 {
	var max float64

	sw.Mutex.RLock()
	defer sw.Mutex.RUnlock()

	for timestamp, window := range sw.Windows {
		if timestamp >= now.Unix()-10 {
			if window.Value > max {
				max = window.Value
			}
		}
	}
	return max
}

// Avg returns the average value over the windows in last 10 seconds.
func (sw *SlidingWindow) Avg(now time.Time) float64 {
	return sw.Sum(now) / 10
}