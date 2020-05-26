package click

// A Doubler is used to count a number of clicks, with each click counted twice.
type Doubler struct {
	Counter
}

// Total returns the clicks counted times two as an int.
func (d Doubler) Total() int {
	return d.Counter.Total() * 2
}
