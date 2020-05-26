package click

import (
	"bufio"
	"fmt"
	"os"
	conv "strconv"
)

const (
	defaultTimes = -1
)

// The Command structure defines a CLI command to execute the click demo.
type Command struct {
	Times int `arg optional name:"times" help:"The number of times you want to click." default:"-1"`
}

// Run executes the click demo command.
func (cmd Command) Run() error {
	var c, d = new(Counter), new(Doubler) // new() creates a new instance of a structure and gives me a pointer to it
	var m = MultiClicker(c, d)

	if cmd.Times < 0 {
		var times, err = readInTimes()
		if err != nil {
			return err
		}
		cmd.Times = times
	}

	DoClicks(m, cmd.Times)
	fmt.Println("Clicked Counter", c.Total(), "times.")
	fmt.Println("Clicked Doubler", d.Total(), "times.")
	return nil
}

func readInTimes() (int, error) {
	var in = bufio.NewScanner(os.Stdin)

	fmt.Print("How many clicks? ")
	if !in.Scan() {
		return defaultTimes, in.Err()
	}

	var clicks, err = conv.Atoi(in.Text())
	if err != nil {
		return defaultTimes, fmt.Errorf("You need to enter an integer value.")
	}

	return clicks, nil
}
