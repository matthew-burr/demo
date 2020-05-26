package click_test

import (
	. "demo/click"
	"fmt"
)

func ExampleDoClicks() {
	var c Counter
	DoClicks(&c, 5)
	fmt.Println("You clicked", c.Total(), "times.")
	// Output: You clicked 5 times.
}
