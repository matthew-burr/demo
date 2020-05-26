package click_test

import (
	. "demo/click"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubler_Click(t *testing.T) {
	tt := []struct {
		name   string
		clicks int
		want   Doubler
	}{
		{"zero clicks", 0, Doubler{Counter: 0}},
		{"one click", 1, Doubler{Counter: 1}},
		{"two clicks", 2, Doubler{Counter: 2}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var d Doubler
			for i := 0; i < tc.clicks; i++ {
				d.Click()
			}
			assert.Equal(t, tc.want, d)
		})
	}
}

func TestDoubler_Total(t *testing.T) {
	tt := []struct {
		name   string
		clicks int
		want   int
	}{
		{"zero clicks", 0, 0},
		{"one click", 1, 2},
		{"two clicks", 2, 4},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var d Doubler
			for i := 0; i < tc.clicks; i++ {
				d.Click()
			}
			var got = d.Total()
			assert.Equal(t, tc.want, got)
		})
	}

}
