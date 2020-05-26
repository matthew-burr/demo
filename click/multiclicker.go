package click

// A ClickerFunc is a function with a signature similar to that of a Clicker.Click.
type ClickerFunc func()

// Click implements the Clicker interface by calling the function in response to a Click.
func (f ClickerFunc) Click() {
	f()
}

// A MultiClicker composes zero or more Clickers into a single Clicker, clicking each one
// in response to a single Click.
func MultiClicker(clickers ...Clicker) ClickerFunc {
	return func() {
		for _, c := range clickers {
			c.Click()
		}
	}
}
