package main

import (
	"fmt"

	"github.com/magefile/mage/sh"

	"github.com/wasilibs/magefiles" // mage:import
)

func init() {
	magefiles.SetLibraryName("protoc-gen-grpc-java")
	magefiles.SetLibraryRepo("grpc/grpc-java")
}

func Snapshot() error {
	return sh.RunV("go", "run", fmt.Sprintf("github.com/goreleaser/goreleaser@%s", verGoReleaser), "release", "--snapshot", "--clean")
}

func Release() error {
	return sh.RunV("go", "run", fmt.Sprintf("github.com/goreleaser/goreleaser@%s", verGoReleaser), "release", "--clean")
}
