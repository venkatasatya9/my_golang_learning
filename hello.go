package main;

import (
		"fmt"
		"satya.com/user/hello/morestrings"
		"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Printf(morestrings.ReverseRunes("!oG, olleH"))
	fmt.Printf(cmp.Diff("Hello World", "Hello Go"))
}