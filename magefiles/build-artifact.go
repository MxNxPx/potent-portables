//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/mg"
)

type Build mg.Namespace

// Create package - aka 'mage b'
func (Build) Build() {
	Build.ZarfVersion(Build{})
	Build.ZarfBuild(Build{})
}

// Output Zarf version
// (sub-Target of 'mage build')
func (Build) ZarfVersion() error {
	return zarf("version")
}

// Create package using Zarf
// (sub-Target of 'mage build')
func (Build) ZarfBuild() error {
	return zarf("package", "create", "--confirm", "--output", "./app", "./app")
}
