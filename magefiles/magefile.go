//go:build mage
// +build mage

package main

var (
	// Aliases are mage aliases of targets
	Aliases = map[string]interface{}{
		"build":  Build.Build,
		"b":      Build.Build,
		"deploy": Deploy.Deploy,
		"d":      Deploy.Deploy,
	}
)

// A var named Default indicates which target is the default.  If there is no
// default, running mage will list the targets available.
//var Default = NA
