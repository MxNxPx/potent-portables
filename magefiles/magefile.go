//go:build mage
// +build mage

package main

var (
	// Aliases are mage aliases of targets
	Aliases = map[string]interface{}{
		"airgap":        Airgap.All,
		"a":             Airgap.All,
		"airgap:deploy": Airgap.ZarfDeploy,
		"airgap:init":   Airgap.ZarfInit,
		"build":         Build.All,
		"b":             Build.All,
		"deploy":        Deploy.All,
		"d":             Deploy.All,
		"deploy:local":  Deploy.ZarfDeploy,
		"deploy:oci":    Deploy.ZarfDeployOCI,
	}
)

// A var named Default indicates which target is the default.  If there is no
// default, running mage will list the targets available.
//var Default = NA
