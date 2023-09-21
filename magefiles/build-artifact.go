//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
)

type Build mg.Namespace

// Create package
func (Build) Build() {
	Build.ZarfVersion(Build{})
	Build.ZarfBuild(Build{})
}

// Output Zarf version
func (Build) ZarfVersion() error {
	return zarf("version")
}

// Create package using Zarf
func (Build) ZarfBuild() error {
	return zarf("package", "create", "--confirm", "--output", "./app", "./app")
}
