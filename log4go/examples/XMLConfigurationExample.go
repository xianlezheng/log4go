package main

import (
	l4g "code.google.com/p/log4go"
	"time"
	"fmt"
)

func main() {
	// Load the configuration (isn't this easy?)
	l4g.LoadConfiguration("/Users/zhengxianle/go/src/code.google.com/p/log4go/examples/example.xml")

	// And now we're ready!
	l4g.Finest("This will only go to those of you really cool UDP kids!  If you change enabled=true.")
	l4g.Debug("Oh no!  %d + %d = %d!", 2, 2, 2+2)
	l4g.Info("About that time, eh chaps?")

	defer func() {
		time.Sleep(time.Second * 10)
		fmt.Println("Stoped")
	}()
}
