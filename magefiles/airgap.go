//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
)

type Airgap mg.Namespace

// Airgap - (aka 'mage a').
func (Airgap) All() {
	Airgap.ZarfInit(Airgap{})

	mg.Deps(Airgap.ZarfInit)
	Airgap.ZarfDeploy(Airgap{})
}

// Airgap Init Cluster - (aka 'mage airgap:init').
func (Airgap) ZarfInit() error {
	return zarf("init", "--components=k3s", "--confirm")
}

// Airgap Deploy - (aka 'mage airgap:deploy').
func (Airgap) ZarfDeploy() error {
	os.Chdir("./app")
	newDir, err := os.Getwd()
	if err != nil {
	}
	fmt.Printf("Current Working Directory: %s\n", newDir)
	filenamePattern := "zarf-package-*.tar.zst" // Change this to your desired wildcard pattern

	filename, err := findFirstFileWithWildcard("./", filenamePattern)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	return zarf("package", "deploy", "--confirm", filename)
}
