// Composing structures
package main

import (
	"demo/click"
	"demo/parallel"

	"github.com/alecthomas/kong"
)

var cli struct {
	Click    click.Command    `cmd help:"Run the 'click' example to demonstrate core features."`
	Parallel parallel.Command `cmd help:"Run the 'parallel' example to demonstrate parallelization features."`
}

// Let's see how this works
func main() {
	var ctx = kong.Parse(&cli)
	ctx.FatalIfErrorf(ctx.Run())
}
