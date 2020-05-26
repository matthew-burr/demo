package click

// A Counter is used to count a number of clicks.
type Counter int

// Click increments the Counter.
func (c *Counter) Click() {
	*c++
}

// Total returns the total number of clicks counted as an int.
func (c *Counter) Total() int {
	return int(*c)
}
