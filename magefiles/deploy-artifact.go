//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
)

type Deploy mg.Namespace

// Install package - aka 'mage d'
// (For existing OCI package, use: `mage deploy oci://pkg-url-here`, or local, use: `mage deploy local`)
func (Deploy) Deploy(ociFlag string) {

	if ociFlag == "local" {
		fmt.Println("No value provided for --oci flag, calling ZarfDeploy")
		Deploy.ZarfDeploy(Deploy{})
	} else {
		Deploy.ZarfDeployOCI(Deploy{}, ociFlag)
	}
}

// Install package using Zarf
// (conditional sub-Target of 'mage deploy')
func (Deploy) ZarfDeploy() error {
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

// Install OCI package using Zarf
// (conditional sub-Target of 'mage deploy')
func (Deploy) ZarfDeployOCI(ociFlag string) error {
	os.Chdir("./app")
	newDir, err := os.Getwd()
	if err != nil {
	}
	fmt.Printf("Current Working Directory: %s\n", newDir)

	return zarf("package", "deploy", ociFlag, "--oci-concurrency=15", "--confirm")
}
