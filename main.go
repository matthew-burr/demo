// Composing structures
package main

import (
	. "demo/click"
	"fmt"
)

// Let's see how this works
func main() {
	var c, d = new(Counter), new(Doubler) // new() creates a new instance of a structure and gives me a pointer to it
	var m = MultiClicker(c, d)
	DoClicks(m, 5)
	fmt.Printf("Counter: %d\nDoubler: %d\n", c.Total(), d.Total())
}
