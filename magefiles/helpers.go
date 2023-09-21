//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var zarf = sh.RunCmd("zarf")

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
