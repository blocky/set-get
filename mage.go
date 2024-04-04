//go:build mage
// +build mage

package main

import (
	"fmt"
	"path"

	"github.com/magefile/mage/sh"
)

var targetDir = "dist"

func build(env map[string]string, target string) {
	sh.RunWithV(env, "go", "build", "-o", target, ".")
}

func BuildAll() {
	for _, os := range []string{"linux", "darwin"} {
		for _, arch := range []string{"amd64", "arm64"} {
			fmt.Printf("Building for %v %v\n", os, arch)
			env := map[string]string{
				"GOOS":   os,
				"GOARCH": arch,
			}

			targetFileName := fmt.Sprintf("set-get-%v-%v", os, arch)
			target := path.Join(targetDir, targetFileName)

			build(env, target)
		}
	}
}

func Build() {
	target := path.Join(targetDir, "set-get")
	build(nil, target)
}

func Clean() {
	sh.Rm(targetDir)
}
