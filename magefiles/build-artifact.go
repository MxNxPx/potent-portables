//go:build mage
// +build mage

// A comment on the package will be output when you list the targets of a
// magefile.

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var zarf = sh.RunCmd("zarf")

// Zarf version
func Zarf() error {
	return zarf("versions")
}

// init example
// func init() {
// 	return ""
// }

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
