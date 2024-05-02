package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "protoc-gen-grpc-java",
		LibraryRepo: "grpc/grpc-java",
		GoReleaser:  true,
	})
	boot.Main()
}
