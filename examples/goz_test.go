package goz

import (
	"fmt"

	"github.com/o289697/goz"
)

func ExampleNewClient() {
	cli := goz.NewClient()

	fmt.Printf("%T", cli)
	// Output: *goz.Request
}
