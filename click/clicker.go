// The click package provides an example of several of Go's core language features.
//
// What's Included
//
// The click package includes examples of the following:
//
// -Derived Types
//
// -Structures and Embedded Types
//
// -Function Types
//
// -Interfaces
//
// -Anonymous Functions
//
// -Closures
//
// -Variadic Functions
//
// -Pointers, Slices, and Variables
//
// -Iteration
//
// -Tests, Examples, and Benchmarks
package click

// A Clicker is any type that receives the Click function.
type Clicker interface {
	Click()
}

// DoClicks will click a Clicker a specified number of times.
func DoClicks(c Clicker, times int) {
	for i := 0; i < times; i++ {
		c.Click()
	}
}
