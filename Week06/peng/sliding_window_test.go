package peng
import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestNewSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		want *SlidingWindow
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSlidingWindow(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlidingWindow_Avg(t *testing.T) {
	type fields struct {
		Windows map[int64]*window
		Mutex   *sync.RWMutex
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := &SlidingWindow{
				Windows: tt.fields.Windows,
				Mutex:   tt.fields.Mutex,
			}
			if got := sw.Avg(tt.args.now); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlidingWindow_Increment(t *testing.T) {
	type fields struct {
		Windows map[int64]*window
		Mutex   *sync.RWMutex
	}
	type args struct {
		i float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := &SlidingWindow{
				Windows: tt.fields.Windows,
				Mutex:   tt.fields.Mutex,
			}
		})
	}
}

func TestSlidingWindow_Max(t *testing.T) {
	type fields struct {
		Windows map[int64]*window
		Mutex   *sync.RWMutex
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := &SlidingWindow{
				Windows: tt.fields.Windows,
				Mutex:   tt.fields.Mutex,
			}
			if got := sw.Max(tt.args.now); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlidingWindow_Sum(t *testing.T) {
	type fields struct {
		Windows map[int64]*window
		Mutex   *sync.RWMutex
	}
	type args struct {
		now time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := &SlidingWindow{
				Windows: tt.fields.Windows,
				Mutex:   tt.fields.Mutex,
			}
			if got := sw.Sum(tt.args.now); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlidingWindow_UpdateMax(t *testing.T) {
	type fields struct {
		Windows map[int64]*window
		Mutex   *sync.RWMutex
	}
	type args struct {
		n float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := &SlidingWindow{
				Windows: tt.fields.Windows,
				Mutex:   tt.fields.Mutex,
			}
		})
	}
}

func TestSlidingWindow_getCurrentWindow(t *testing.T) {
	type fields struct {
		Windows map[int64]*window
		Mutex   *sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   *window
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := &SlidingWindow{
				Windows: tt.fields.Windows,
				Mutex:   tt.fields.Mutex,
			}
			if got := sw.getCurrentWindow(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCurrentWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlidingWindow_removeOldWindows(t *testing.T) {
	type fields struct {
		Windows map[int64]*window
		Mutex   *sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sw := &SlidingWindow{
				Windows: tt.fields.Windows,
				Mutex:   tt.fields.Mutex,
			}
		})
	}
}