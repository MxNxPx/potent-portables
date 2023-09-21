//go:build mage
// +build mage

// A comment on the package will be output when you list the targets of a
// magefile.

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	// Aliases are mage aliases of targets
	Aliases = map[string]interface{}{
		"build":  Build.Build,
		"deploy": Deploy.Deploy,
	}
)

var zarf = sh.RunCmd("zarf")

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

type Deploy mg.Namespace

// Install package
// If deploying using an existing OCI package, use: `mage deploy oci://pkg-url-here`, or deploy from local, use: `mage deploy local`
func (Deploy) Deploy(ociFlag string) {

	if ociFlag == "local" {
		fmt.Println("No value provided for --oci flag, calling ZarfDeploy")
		Deploy.ZarfDeploy(Deploy{})
	} else {
		Deploy.ZarfDeployOCI(Deploy{}, ociFlag)
	}
}

// Install package using Zarf
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
func (Deploy) ZarfDeployOCI(ociFlag string) error {
	os.Chdir("./app")
	newDir, err := os.Getwd()
	if err != nil {
	}
	fmt.Printf("Current Working Directory: %s\n", newDir)

	return zarf("package", "deploy", ociFlag, "--oci-concurrency=15", "--confirm")
}

func runWith(env map[string]string, cmd string, inArgs ...any) error {
	s := argsToStrings(inArgs...)
	return sh.RunWith(env, cmd, s...)
}

func runCmd(env map[string]string, cmd string, args ...any) error {
	if mg.Verbose() {
		return runWith(env, cmd, args...)
	}
	output, err := sh.OutputWith(env, cmd, argsToStrings(args...)...)
	if err != nil {
		fmt.Fprint(os.Stderr, output)
	}

	return err
}

func argsToStrings(v ...any) []string {
	var args []string
	for _, arg := range v {
		switch v := arg.(type) {
		case string:
			if v != "" {
				args = append(args, v)
			}
		case []string:
			if v != nil {
				args = append(args, v...)
			}
		default:
			panic("invalid type")
		}
	}

	return args
}

func findFirstFileWithWildcard(dir, wildcard string) (string, error) {
	// Use filepath.Glob to list files that match the wildcard in the given directory
	matches, err := filepath.Glob(filepath.Join(dir, wildcard))
	if err != nil {
		return "", err
	}

	// Check if there are any matches
	if len(matches) == 0 {
		return "", fmt.Errorf("No matching files found")
	}

	// Extract the first matching filename without the directory
	filename := filepath.Base(matches[0])
	return filename, nil
}
