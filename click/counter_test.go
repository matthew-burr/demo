package click_test

import (
	. "demo/click"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter_Click(t *testing.T) {
	tt := []struct {
		name   string
		clicks int
		want   Counter
	}{
		{"zero clicks", 0, Counter(0)},
		{"one click", 1, Counter(1)},
		{"two clicks", 2, Counter(2)},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var c Counter
			for i := 0; i < tc.clicks; i++ {
				c.Click()
			}
			assert.Equal(t, tc.want, c)
		})
	}
}

func TestCounter_Total(t *testing.T) {
	tt := []struct {
		name   string
		clicks int
		want   int
	}{
		{"zero clicks", 0, 0},
		{"one click", 1, 1},
		{"two clicks", 2, 2},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var c Counter
			for i := 0; i < tc.clicks; i++ {
				c.Click()
			}
			var got = c.Total()
			assert.Equal(t, tc.want, got)
		})
	}

}
