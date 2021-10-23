package main

// Visit pkg.go.dev and search for a "quote" package.
// Locate and click the rsc.io/quote package in search results

import (
	"fmt"

	"rsc.io/quote"
)

// run command " go mod tidy"
func main() {
	fmt.Print(quote.Go())
}
