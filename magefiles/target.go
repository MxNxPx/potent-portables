//go:build mage
// +build mage

// A comment on the package will be output when you list the targets of a
// magefile.
package main

import (
    "log"
)

// The first sentence in the comment will be the short help text shown with mage -l.
// The rest of the comment is long help text that will be shown with mage -h <target>
func Hello() {
    // by default, the log stdlib package will be set to discard output.
    // Running with mage -v will set the output to stdout.
    log.Printf("Hi!")
}

