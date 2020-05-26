package click_test

import (
	. "demo/click"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiClicker_Click(t *testing.T) {
	var x, y Counter
	var m = MultiClicker(&x, &y)
	m.Click()
	m.Click()
	var want = Counter(2)
	assert.Equal(t, want, x)
	assert.Equal(t, want, y)
}

func BenchmarkMultiClicker_Click(b *testing.B) {
	bb := []struct {
		name string
		size int
	}{
		{"little", 10},
		{"medium", 100},
		{"big", 10000},
	}

	for _, bc := range bb {
		var counters = make([]Clicker, bc.size)
		for i := range counters {
			counters[i] = new(Counter)
		}
		var m = MultiClicker(counters...)
		b.Run(bc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				m.Click()
			}
		})
	}
}

func BenchmarkMultiClicker_Click_Profile(b *testing.B) {
	var counters = make([]Clicker, 10000)
	for i := range counters {
		counters[i] = new(Counter)
	}
	var m = MultiClicker(counters...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Click()
	}
}
